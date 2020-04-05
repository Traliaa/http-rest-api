package apiserver

import (
	"database/sql"
	"fmt"
	"github.com/Traliaa/http-rest-api/internal/app/store/sqlstore"
	"github.com/gorilla/sessions"
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

	return http.ListenAndServe(config.BindAddr, srv)
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
