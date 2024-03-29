package main

import (
	_ "database/sql"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"knock-knock/controllers"
	"knock-knock/driver"
	"knock-knock/modelLoading"
	"log"
	"net/http"
	"os"
)

/*type ViewDataUsers struct {
	Title   string `json:"title"`
	Content []models.User `json:"content"`
}
type ViewDataUser struct {
	Title   string `json:"title"`
	Content models.User   `json:"content"`
}*/

func main() {
	var db *sqlx.DB
	controller := controllers.Controller{}
	db = driver.ConnectDB()
	router := mux.NewRouter()

	modelLoading.LoadUsers(db)
	modelLoading.LoadTeams(db)
	modelLoading.LoadMarks(db)

	log.Println("Users and Teams are loaded")

	router.HandleFunc("/signup", controller.Signup(db)).Methods("POST")
	router.HandleFunc("/signin", controller.Signin(db)).Methods("POST")
	router.HandleFunc("/users", controller.GetUsers(db)).Methods("GET")
	router.HandleFunc("/users/{id}", controller.GetUser(db)).Methods("GET")
	router.HandleFunc("/logout", controller.Logout()).Methods("POST")
	router.HandleFunc("/mypage", controller.GetMyPage(db)).Methods("GET")
	router.HandleFunc("/teams", controller.GetTeams(db)).Methods("GET")
	//router.HandleFunc("/teams/", controller.GetTeams(db)).Methods("GET")
	router.HandleFunc("/teams/{id}", controller.GetTeam(db)).Methods("GET")
	router.HandleFunc("/teamList", controller.GetTeamList(db)).Methods("GET")

	fmt.Println("Server is listening...")

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	log.Fatal(http.ListenAndServe(":8181", loggedRouter))
}
