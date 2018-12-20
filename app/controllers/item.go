package controllers

import (
	"database/sql"
	"encoding/json"
	"simple-go-rest-api/app/driver"
	"simple-go-rest-api/app/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Controller : data struct for handle functions
type Controller struct{}

var items []models.Item

func apiLogger(rq *http.Request) {
	log.Println(rq.RemoteAddr, rq.RequestURI, rq.Method)
}

// Controller.GetItems : handle func for GET method (get all items)
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

// Controller.GetItem : handle func for GET method (get single item)
func (c Controller) GetItem(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, rq *http.Request) {
		var item models.Item
		params := mux.Vars(rq)

		row := db.QueryRow("select * from items where id=$1", params["id"])
		err := row.Scan(&item.ID, &item.Name1, &item.Name2, &item.Name3)
		driver.LogFatal(err)

		// defer row.Close() ???
		json.NewEncoder(w).Encode(item)
		apiLogger(rq)
	}
}

// Controller.AddItem : handle func for POST method (add single item)
func (c Controller) AddItem(db *sql.DB) http.HandlerFunc {
 	return func(w http.ResponseWriter, rq *http.Request) {
		var item models.Item
		var itemID int

		json.NewDecoder(rq.Body).Decode(&item)
		//log.Println(item)
		err := db.QueryRow("insert into items (name1, name2, name3) values ($1, $2, $3) RETURNING id", item.Name1, item.Name2, item.Name3).Scan(&itemID)
		driver.LogFatal(err)
		json.NewEncoder(w).Encode(itemID)
		apiLogger(rq)
	}
}

// Controller.UpdateItem : handle func for PUT method (update single item)
func (c Controller) UpdateItem(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, rq *http.Request) {
		var item models.Item

		json.NewDecoder(rq.Body).Decode(&item)

		result, err := db.Exec("update items set name1=$1, name2=$2, name3=$3 where id=$4 RETURNING id", &item.Name1, &item.Name2, &item.Name3, &item.ID)
		driver.LogFatal(err)

		rowsUpdated, err := result.RowsAffected()
		driver.LogFatal(err)

		json.NewEncoder(w).Encode(rowsUpdated)
		apiLogger(rq)
	}
}

//Controller.RemoveItem : handle func for DELETE method (delete single item)
func (c Controller) RemoveItem(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, rq *http.Request) {
		params := mux.Vars(rq)

		result, err := db.Exec("delete from items where id=$1", params["id"])
		driver.LogFatal(err)

		rowsDeleted, err := result.RowsAffected()
		driver.LogFatal(err)

		json.NewEncoder(w).Encode(rowsDeleted)
		apiLogger(rq)
	}
}
