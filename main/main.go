package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"simple-go-rest-api/controllers"
	"simple-go-rest-api/driver"

	"github.com/gorilla/mux"
)

var db *sql.DB

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "5002"
	}
	return ":" + port
}

func main() {

	db = driver.ConnectDB()
	router := mux.NewRouter()
	controller := controllers.Controller{}

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(port(), router))
}
