package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"simple-go-rest-api/driver"
	"simple-go-rest-api/models"

	"github.com/gorilla/mux"
)

// Controller : data struct for handle functions
type Controller struct{}

var books []models.Book

func apiLogger(rq *http.Request) {
	log.Println(rq.RemoteAddr, rq.RequestURI, rq.Method)
}

// GetBooks : handle func for GET method (get all books)
func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, rq *http.Request) {
		var book models.Book
		books = []models.Book{}

		rows, err := db.Query("select * from books")
		driver.LogFatal(err)
		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&book.ID, &book.Name1, &book.Name2, &book.Name3)
			driver.LogFatal(err)
			books = append(books, book)
		}
		json.NewEncoder(w).Encode(books)
		apiLogger(rq)
	}
}

// GetBook : handle func for GET method (get single book)
func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, rq *http.Request) {
		var book models.Book
		params := mux.Vars(rq)

		row := db.QueryRow("select * from books where id=$1", params["id"])
		err := row.Scan(&book.ID, &book.Name1, &book.Name2, &book.Name3)
		driver.LogFatal(err)

		// defer row.Close() ???
		json.NewEncoder(w).Encode(book)
		apiLogger(rq)
	}
}

// AddBook : handle func for POST method (add single book)
func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, rq *http.Request) {
		var book models.Book
		var bookID int

		json.NewDecoder(rq.Body).Decode(&book)
		//log.Println(book)
		err := db.QueryRow("insert into books (name1, name2, name3) values ($1, $2, $3) RETURNING id", book.Name1, book.Name2, book.Name3).Scan(&bookID)
		driver.LogFatal(err)
		json.NewEncoder(w).Encode(bookID)
		apiLogger(rq)
	}
}

// UpdateBook : handle func for PUT method (update single book)
func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, rq *http.Request) {
		var book models.Book

		json.NewDecoder(rq.Body).Decode(&book)

		result, err := db.Exec("update books set name1=$1, name2=$2, name3=$3 where id=$4 RETURNING id", &book.Name1, &book.Name2, &book.Name3, &book.ID)
		driver.LogFatal(err)

		rowsUpdated, err := result.RowsAffected()
		driver.LogFatal(err)

		json.NewEncoder(w).Encode(rowsUpdated)
		apiLogger(rq)
	}
}

// RemoveBook : handle func for DELETE method (delete single book)
func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, rq *http.Request) {
		params := mux.Vars(rq)

		result, err := db.Exec("delete from books where id=$1", params["id"])
		driver.LogFatal(err)

		rowsDeleted, err := result.RowsAffected()
		driver.LogFatal(err)

		json.NewEncoder(w).Encode(rowsDeleted)
		apiLogger(rq)
	}
}
