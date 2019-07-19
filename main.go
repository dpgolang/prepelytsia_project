package main

import (
	_"database/sql"
	"github.com/jmoiron/sqlx"
	"fmt"
	"github.com/gorilla/mux"
	_"github.com/lib/pq"
	"log"
	"net/http"
	"knock-knock/controllers"
	"knock-knock/driver"
	_"knock-knock/models"
	"os"
	"github.com/gorilla/handlers"
)



/*type ViewDataUsers struct {
	Title   string `json:"title"`
	Content []models.User `json:"content"`
}
type ViewDataUser struct {
	Title   string `json:"title"`
	Content models.User   `json:"content"`
}*/

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}


func main() {
	var db *sqlx.DB
	controller := controllers.Controller{}
	db = driver.ConnectDB()
	router := mux.NewRouter()
	router.HandleFunc("/signup", controller.Signup(db)).Methods("POST")
	router.HandleFunc("/signin", controller.Signin(db)).Methods("POST")
	router.HandleFunc("/users", controller.GetUsers(db)).Methods("GET")
	router.HandleFunc("/users/{id}", controller.GetUser(db)).Methods("GET")
	router.HandleFunc("/logout",controller.Logout()).Methods("POST")
	fmt.Println("Server is listening...")

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	log.Fatal(http.ListenAndServe(":8181", loggedRouter))
}
