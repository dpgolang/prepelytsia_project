package driver

import (
	//"database/sql"
	"log"
	"fmt"
	"github.com/subosito/gotenv"
	"os"
    "github.com/jmoiron/sqlx"

	_"github.com/lib/pq"
)

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() *sqlx.DB {
	var err error
	var db *sqlx.DB
	host:= os.Getenv("HOST_POSTGRES")
	port:= os.Getenv("PORT_POSTGRES")
	user:= os.Getenv("USER_POSTGRES")
	password:= os.Getenv("PASSWORD_POSTGRES")
	dbname:= os.Getenv("DBNAME_POSTGRES")
	/*pgURL,err := pq.ParseURL(os.Getenv("POSTGRESQL_URL"))
	logFatal(err)
	log.Println(pgURL)*/
	connStr := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sqlx.Open("postgres", connStr)
	logFatal(err)
	err = db.Ping()
	logFatal(err)
	return db
}
