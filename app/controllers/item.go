package controllers

import (
	"database/sql"
	"encoding/json"
	"app/driver"
	"app/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Controller: data struct for handle functions
type Controller struct{}

var items []models.Item

func apiLogger(rq *http.Request) {
	log.Println(rq.RemoteAddr, rq.RequestURI, rq.Method)
}

// Controller.GetItems: handle func for GET method (get all items)
func (c Controller) GetItems(db *sql.DB) http.HandlerFunc {
}

// Controller.GetItem: handle func for GET method (get single item)
func (c Controller) GetItem(db *sql.DB) http.HandlerFunc {
}

// Controller.AddItem: handle func for POST method (add single item)
func (c Controller) AddItem(db *sql.DB) http.HandlerFunc {
}

// Controller.UpdateItem: handle func for PUT method (update single item)
func (c Controller) UpdateItem(db *sql.DB) http.HandlerFunc {
}

//Controller.RemoveItem: handle func for DELETE method (delete single item)
func (c Controller) RemoveItem(db *sql.DB) http.HandlerFunc {
}
