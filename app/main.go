package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	ALL_BOOKS = "/books"
	ONE_BOOK = "/books/{id}"
	CREATE_BOOK = "/books"

	LIVENESS_PROBE = "/live"
	READINESS_PROBE = "/ready"
)

func main() {
	db := &DatabaseConnection{}

	err := db.OpenConnection()

	if err != nil {
		log.Fatal(fmt.Sprintf("Could not stablish connection with DB: %s", err.Error()))
		return
	}

	booksController := &BookController{
		db: db,
	}

	router := mux.NewRouter()

	router.HandleFunc(ALL_BOOKS, booksController.ListBooks).Methods(http.MethodGet)
	router.HandleFunc(ONE_BOOK, booksController.ListOneBook).Methods(http.MethodGet)
	router.HandleFunc(CREATE_BOOK, booksController.CreateBook).Methods(http.MethodPost)
	router.HandleFunc(ONE_BOOK, booksController.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc(ONE_BOOK, booksController.DeleteBook).Methods(http.MethodDelete)

	router.HandleFunc(LIVENESS_PROBE, livenessProbe).Methods(http.MethodGet)
	router.HandleFunc(READINESS_PROBE, readinessProbe).Methods(http.MethodGet)


	log.Fatal(http.ListenAndServe(":8000", router))
}
