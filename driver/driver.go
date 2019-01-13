package driver

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
	"github.com/subosito/gotenv"
)

var db *sql.DB

// LogFatal func for logging errors
func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// env init func
func pgInit() {
	gotenv.Load()
}

// ConnectDB func for db connection init
func ConnectDB() *sql.DB {
	pgInit()
	pgURL, err := pq.ParseURL(os.Getenv("PG_URL"))
	log.Println("exported PG_URL from .env:", pgURL)
	LogFatal(err)

	pgDB := os.Getenv("PG_DB")
	log.Println("exported PG_DB from .env:", pgDB)

	db, err = sql.Open(pgDB, pgURL)
	LogFatal(err)

	err = db.Ping()
	LogFatal(err)

	return db
}
