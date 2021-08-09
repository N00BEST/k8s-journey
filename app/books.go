package main

import (
	"encoding/json"
	"net/http"
)

const (
	GET_BOOKS_QUERY = "SELECT id, title, author FROM `books`"
	GET_ONE_BOOK_QUERY = "SELECT id, title, author FROM `books` where id = ?"
	CREATE_BOOK_QUERY = "INSERT INTO `books` (title, author) VALUES (?, ?)"
	UPDATE_BOOK_QUERY = "UPDATE `books` SET title = ?, author = ? WHERE id = ?"
	DELETE_BOOK_QUERY = "DELETE FROM `books` WHERE id = ?"
)

type Book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
}

type BookController struct {
	db *DatabaseConnection
}

func (bc *BookController) ListBooks(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode([]Book{
		{
			ID: "1",
			Title: "My book",
			Author: "Me, myself and I",
		},
	})
	return

	//rows, err := bc.db.Connection.Query(GET_BOOKS_QUERY)
	//
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	json.NewEncoder(w).Encode(fmt.Sprintf(`{"error":"%s"}`, err.Error()))
	//	return
	//}
	//
	//books := []Book{}
	//
	//for rows.Next() {
	//	var book Book
	//	err = rows.Scan(&book.ID, &book.Title)
	//	if err != nil {
	//		w.WriteHeader(http.StatusInternalServerError)
	//		json.NewEncoder(w).Encode(fmt.Sprintf(`{"error":"%s"}`, err.Error()))
	//	}
	//	// and append it to the array
	//	books = append(books, book)
	//}
	//
	//json.NewEncoder(w).Encode(books)
}

func (bc *BookController) ListOneBook(w http.ResponseWriter, req *http.Request) {

}

func (bc *BookController) CreateBook(w http.ResponseWriter, req *http.Request) {

}

func (bc *BookController) UpdateBook(w http.ResponseWriter, req *http.Request) {

}

func (bc *BookController) DeleteBook(w http.ResponseWriter, req *http.Request) {

}