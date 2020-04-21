package apiserver

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Traliaa/http-rest-api/internal/app/model"
	"github.com/Traliaa/http-rest-api/internal/app/store"
	"github.com/Traliaa/http-rest-api/internal/app/webserver"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
	"time"
)

const userPass = "nogin:pedik"
const unauth = http.StatusUnauthorized

// Account request model
type Account struct {
	// Id of the account
	ID string `json:"id"`
	// First Name of the account holder
	FirstName string `json:"first_name"`
	// Last Name of the account holder
	LastName string `json:"last_name"`
	// User Name of the account holder
	UserName string `json:"user_name"`
}

// Account response payload
// swagger:response accountRes
type swaggAccountRes struct {
	// in:body
	Body Account
}

// Error Bad Request
// swagger:response badReq
type swaggReqBadRequest struct {
	// in:body
	Body struct {
		// HTTP status code 400 -  Bad Request
		Code int `json:"code"`
	}
}

// Error Not Found
// swagger:response notFoundReq
type swaggReqNotFound struct {
	// in:body
	Body struct {
		// HTTP status code 404 -  Not Found
		Code int `json:"code"`
	}
}

const (
	sessionName        = "AmiCorp"
	ctxKeyUser  ctxKey = iota
	ctxKeyRequestID
)

var (
	errIncorrectEmailOrPassaword = errors.New("Incorrect Email or Password")
	errNotAuthenticated          = errors.New("Not Authenticated")
	upgrader                     = websocket.Upgrader{}
)

type ctxKey int32

type server struct {
	router       *mux.Router
	logger       *logrus.Logger
	store        store.Store
	sessionStore sessions.Store
}

func NewServer(store store.Store, sessionStore sessions.Store) *server {
	s := &server{
		router:       mux.NewRouter(),
		logger:       logrus.New(),
		store:        store,
		sessionStore: sessionStore,
	}
	s.configureRouter()
	return s

}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)

}

func (s *server) configureRouter() {
	s.router.Handle("/metrics", promhttp.Handler())
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("src/swager"))))
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
	s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods("POST")
	s.router.HandleFunc("/echo", s.echo)
	s.router.HandleFunc("/", s.home)
	private := s.router.PathPrefix("/private").Subrouter()
	private.Use(s.authenticateUsers)
	private.HandleFunc("/whoami", s.handleWhoami()).Methods("GET")

}

func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(ctxKeyRequestID),
		})
		logger.Infof("Started %s %s", r.Method, r.RequestURI)
		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(w, r)
		logger.Infof(
			"Completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Now().Sub(start))

	})
}

func (s *server) authenticateUsers(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		id, ok := session.Values["user_id"]
		if !ok {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}
		u, err := s.store.User().Find(id.(int))
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
	})
}

// swagger:operation GET /accounts/{id} accounts getAccount
// ---
// summary: Return an Account provided by the id.
// description: If the account is found, account will be returned else Error Not Found (404) will be returned.
// parameters:
// - name: id
//   in: path
//   description: id of the account
//   type: string
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/accountRes"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFoundReq"
func (s *server) handleWhoami() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*model.User))
	}

}

func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
	})

}

// swagger:operation POST /users/{email} accounts getAccount
// ---
// summary: Return an Account provided by the id.
// description: If the account is found, account will be returned else Error Not Found (404) will be returned.
// parameters:
// - name: email
//   in: path
//   description: id of the account
//   type: string
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/accountRes"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFoundReq"
func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			logrus.Error(err)

			return
		}

		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}
		if err := s.store.User().Create(u); err != nil {
			fmt.Print("er")
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)
	}
}

// swagger:operation POST /accounts/{id} accounts getAccount
// ---
// summary: Return an Account provided by the id.
// description: If the account is found, account will be returned else Error Not Found (404) will be returned.
// parameters:
// - name: id
//   in: path
//   description: id of the account
//   type: string
//   required: true
// responses:
//   "200":
//     "$ref": "#/responses/accountRes"
//   "400":
//     "$ref": "#/responses/badReq"
//   "404":
//     "$ref": "#/responses/notFoundReq"
func (s *server) handleSessionsCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			logrus.Error(err)

			return
		}
		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassaword)
			return
		}

		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		session.Values["user_id"] = u.ID
		if err := s.sessionStore.Save(r, w, session); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *server) home(w http.ResponseWriter, r *http.Request) {

	webserver.HomeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

func (s *server) echo(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if !strings.HasPrefix(auth, "Basic ") {
		logrus.Print("Invalid authorization:", auth)
		http.Error(w, http.StatusText(unauth), unauth)
		return
	}
	up, err := base64.StdEncoding.DecodeString(auth[6:])
	if err != nil {
		logrus.Print("authorization decode error:", err)
		http.Error(w, http.StatusText(unauth), unauth)
		return
	}
	if string(up) != userPass {
		logrus.Print("invalid username:password:", string(up))
		http.Error(w, http.StatusText(unauth), unauth)
		return
	}
	io.WriteString(w, "Goodbye, World!")

	c, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)

	if err != nil {
		logrus.Error("upgrade:", err)
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}
	defer c.Close()
	webserver.SendClient(c)

}
