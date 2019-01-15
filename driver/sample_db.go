package driver

import (
	"database/sql"
	"simple-go-rest-api/models"
)

var sampleBooks = []models.Book{
	{1, "Golang1", "Golangovsky", "1996"},
	{2, "Golang2", "Golangovitz", "1997"},
	{3, "Golang3", "Golangoff", "1998"},
	{4, "Golang4", "Golangovich", "1999"},
}

func initValuesInDB(db *sql.DB) {
	var bookID int

	for _, item := range sampleBooks {
		err := db.QueryRow("insert into books (title, author, year) values ($1, $2, $3) RETURNING id", item.Title, item.Author, item.Year).Scan(&bookID)
		LogFatal(err)
	}
}
