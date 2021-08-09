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

	router.HandleFunc(ALL_BOOKS, booksController.ListBooks).Methods("GET")
	router.HandleFunc(ONE_BOOK, booksController.ListOneBook).Methods("GET")
	router.HandleFunc(CREATE_BOOK, booksController.CreateBook).Methods("POST")
	router.HandleFunc(ONE_BOOK, booksController.UpdateBook).Methods("PUT")
	router.HandleFunc(ONE_BOOK, booksController.DeleteBook).Methods("DELETE")


	log.Fatal(http.ListenAndServe(":8081", router))
}
