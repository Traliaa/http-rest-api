package apiserver

import (
	"database/sql"
	"fmt"
	"github.com/Traliaa/http-rest-api/internal/app/monitoring"
	api "github.com/Traliaa/http-rest-api/internal/app/proto"
	"github.com/Traliaa/http-rest-api/internal/app/store/sqlstore"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURl)
	if err != nil {
		return err
	}
	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := NewServer(store, sessionStore)

	//Start monitoring
	monitoring.RecordMetrics()

	return http.ListenAndServe(config.BindAddr, srv)
}
func StartGRPCServer() {
	//db, err := newDB(config.DatabaseURl)
	//if err != nil {
	//	return err
	//}
	//store := sqlstore.New(db)
	//sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	//srv := NewServer(store, sessionStore)

	//Start monitoring
	monitoring.RecordMetrics()

	//return http.ListenAndServe(config.BindAddr, srv)
	s := grpc.NewServer()
	srv := &GRPCServer{}
	api.RegisterLoginServer(s, srv)
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		logrus.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		logrus.Fatal(err)
	}

}

func newDB(databaseURl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURl)
	if err != nil {
		fmt.Print("no open")
		return nil, err
	}
	//defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Print("no ping")
		return nil, err

	}
	return db, nil
}
