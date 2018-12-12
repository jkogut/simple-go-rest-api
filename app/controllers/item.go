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
	return func(w http.ResponseWriter, rq *http.Request) {
		var item models.Item
		items = []models.Item{}

		rows, err := db.Query("select * from items")
		driver.LogFatal(err)
		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&item.ID, &item.Name1, &item.Name2, &item.Name3)
			driver.LogFatal(err)
			items = append(items, item)
		}
		json.NewEncoder(w).Encode(items)
		apiLogger(rq)
	}
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