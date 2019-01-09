package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"simple-go-rest-api/controllers"
	"simple-go-rest-api/driver"
	"simple-go-rest-api/models"

	"github.com/gorilla/mux"
)

var items []models.Item
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

	router.HandleFunc("/items", controller.GetItems(db)).Methods("GET")
	router.HandleFunc("/items/{id}", controller.GetItem(db)).Methods("GET")
	router.HandleFunc("/items", controller.AddItem(db)).Methods("POST")
	router.HandleFunc("/items", controller.UpdateItem(db)).Methods("PUT")
	router.HandleFunc("/items/{id}", controller.RemoveItem(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(port(), router))
}
