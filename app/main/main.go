package main

import (
	"database/sql"
	"app/controllers"
	"app/driver"
	"app/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var items []models.Item
var db *sql.DB

func main() {

	db = driver.ConnectDB()
	router := mux.NewRouter()
	controller := controllers.Controller{}

	router.HandleFunc("/items", controller.GetItems(db)).Methods("GET")
	router.HandleFunc("/items/{id}", controller.GetItem(db)).Methods("GET")
	router.HandleFunc("/items", controller.AddItem(db)).Methods("POST")
	router.HandleFunc("/items", controller.UpdateItem(db)).Methods("PUT")
	router.HandleFunc("/items/{id}", controller.RemoveItem(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5002", router))
}
