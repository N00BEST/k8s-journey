package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

const (
	GET_BOOKS_QUERY = "SELECT id, title, author FROM `books`"
	GET_ONE_BOOK_QUERY = "SELECT id, title, author FROM `books` where id = ?"
	CREATE_BOOK_QUERY = "INSERT INTO `books` (title, author) VALUES (?, ?)"
	UPDATE_BOOK_QUERY = "UPDATE `books` SET title = ?, author = ? WHERE id = ?"
	DELETE_BOOK_QUERY = "DELETE FROM `books` WHERE id = ?"
)

type Book struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
}

type BookController struct {
	db *DatabaseConnection
}

func (bc *BookController) ListBooks(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := bc.db.Connection.Query(GET_BOOKS_QUERY)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(fmt.Sprintf(`{"error":"%s"}`, err.Error()))
		return
	}

	books := []Book{}

	for rows.Next() {
		var book Book
		err = rows.Scan(&book.ID, &book.Title, &book.Author)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(fmt.Sprintf(`{"error":"%s"}`, err.Error()))
			return
		}
		// and append it to the array
		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)
}

func (bc *BookController) ListOneBook(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	queryId := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(queryId, 0, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(fmt.Sprintf(`{"error":"%s"}`, err.Error()))
		return
	}

	var book Book

	row := bc.db.Connection.QueryRow(GET_ONE_BOOK_QUERY, id)

	err = row.Scan(&book.ID, &book.Title, &book.Author)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(fmt.Sprintf(`{"error":"%s"}`, err.Error()))
		return
	}

	if book.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func (bc *BookController) CreateBook(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book
	err := json.NewDecoder(req.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(fmt.Sprintf(`{"error":"%s"}`, err.Error()))
	} else {
		result, err := bc.db.Connection.Exec(CREATE_BOOK_QUERY, book.Title, book.Author)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(fmt.Sprintf(`{"error":"%s"}`, err.Error()))
		} else {
			id, err := result.LastInsertId()

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(fmt.Sprintf(`{"error":"%s"}`, err.Error()))
				return
			}

			book.ID = id
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(book)
		}
	}
}

func (bc *BookController) UpdateBook(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	queryId := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(queryId, 0, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(fmt.Sprintf(`{"error":"%s"}`, err.Error()))
		return
	} else {
		var book Book
		err := json.NewDecoder(req.Body).Decode(&book)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(fmt.Sprintf(`{"error":"%s"}`, err.Error()))
			return
		} else {
			result, err := bc.db.Connection.Exec(UPDATE_BOOK_QUERY, book.Title, book.Author, id)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(fmt.Sprintf(`{"error":"%s"}`, err.Error()))
				return
			} else {
				affectedRows, err := result.RowsAffected()

				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode(fmt.Sprintf(`{"error":"%s"}`, err.Error()))
					return
				} else {
					if affectedRows == 0 {
						w.WriteHeader(http.StatusNotFound)
						return
					} else {
						book.ID = id
						json.NewEncoder(w).Encode(book)
						return
					}
				}
			}
		}
	}
}

func (bc *BookController) DeleteBook(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	queryId := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(queryId, 0, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(fmt.Sprintf(`{"error":"%s"}`, err.Error()))
		return
	} else {
		result, err := bc.db.Connection.Exec(DELETE_BOOK_QUERY, id)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(fmt.Sprintf(`{"error":"%s"}`, err.Error()))
			return
		} else {
			affectedRows, err := result.RowsAffected()

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(fmt.Sprintf(`{"error":"%s"}`, err.Error()))
				return
			} else {
				if affectedRows == 0 {
					w.WriteHeader(http.StatusNotFound)
					return
				} else {
					w.WriteHeader(http.StatusOK)
					return
				}
			}
		}
	}
}