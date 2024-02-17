package main

import (
	"beer-api/core/beer"
	"beer-api/web/handlers"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {
	db, err := sql.Open("sqlite3", "data/beer.db")

	if err != nil {
		log.Fatal(err) // vai fazer log e interromper o programa
	}
	defer db.Close()

	service := beer.NewService(db)

	r := mux.NewRouter()
	// handlers
	n := negroni.New(
		negroni.NewLogger(),
	)

	// handlers
	handlers.MakeHandlers(service, r, n)
	http.Handle("/", r)

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)

	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":4000",
		Handler:      http.DefaultServeMux,
		ErrorLog:     logger,
	}

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
