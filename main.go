package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	//"github.com/subosito/gotenv"
	"knock-knock/controllers"
	"knock-knock/driver"
	"knock-knock/models"
	//_"knock-knock/repository"
	_ "os"
)

var users []models.User
var db *sql.DB

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
	/* pgURL,err := pq.ParseURL(os.Getenv("POSTGRESQL_URL"))
	   logFatal(err)
	   log.Println(pgURL)*/
	//http.HandleFunc("/signin", Signin)
	controller := controllers.Controller{}
	db = driver.ConnectDB()
	router := mux.NewRouter()
	router.HandleFunc("/signup", controller.Signup(db)).Methods("POST")
	router.HandleFunc("/signin", controller.Signin(db)).Methods("POST")
	router.HandleFunc("/users", controller.GetUsers(db)).Methods("GET")
	router.HandleFunc("/users/{id}", controller.GetUser(db)).Methods("GET")
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", router)
}
