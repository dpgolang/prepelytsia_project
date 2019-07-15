package driver

import (
	"database/sql"
	"log"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123123"
	dbname   = "knock-knock"
)

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func ConnectDB() *sql.DB {
	var err error
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", connStr)
	logFatal(err)

	err = db.Ping()
	logFatal(err)
	return db
}
