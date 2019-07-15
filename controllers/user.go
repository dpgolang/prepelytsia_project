package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"knock-knock/models"
	"knock-knock/repository/user"
	"log"
	"net/http"
	"strconv"
)

var users []models.User

type Controller struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) GetUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		users = []models.User{}

		userRepo := userRepository.UserRepository{}
		users = userRepo.GetUsers(db, user, users)
		data := struct {
			Title   string
			Content []models.User
		}{
			Title:   "Users",
			Content: users,
		}
		tmpl, _ := template.ParseFiles("templates/users.html")
		tmpl.Execute(w, data)
		json.NewEncoder(w).Encode(users)
	}
}

func (c Controller) GetUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var user models.User
		userRepo := userRepository.UserRepository{}
		id, err := strconv.Atoi(params["id"])
		logFatal(err)
		user = userRepo.GetUser(db, user, id)
		if user.Id == 0 {
			data := struct {
				Title   string
				Content models.User
			}{
				Title:   "There is no user with such id!",
				Content: user,
			}
			tmpl, _ := template.ParseFiles("templates/user.html")
			tmpl.Execute(w, data)
			json.NewEncoder(w).Encode(user)
			return
		}
		logFatal(err)
		data := struct {
			Title   string
			Content models.User
		}{
			Title:   user.Firstname + " " + user.Lastname,
			Content: user,
		}
		tmpl, _ := template.ParseFiles("templates/user.html")
		tmpl.Execute(w, data)
		json.NewEncoder(w).Encode(user)
	}
}

func (c Controller) Signup(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		var userID int

		err := json.NewDecoder(r.Body).Decode(&user)
		logFatal(err)
		/*if err != nil {
			// If there is something wrong with the request body, return a 400 status
			w.WriteHeader(http.StatusBadRequest)
			return
		}*/
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
		user.Password = string(hashedPassword)
		userRepo := userRepository.UserRepository{}
		userID = userRepo.Signup(db, user)
		json.NewEncoder(w).Encode(&userID)
	}
}
func (c Controller) Signin(db *sql.DB) http.HandlerFunc { //как отлавливать когда айди не найдет
	return func(w http.ResponseWriter, r *http.Request) {
		userFromBase := models.User{}
		userChecking := models.User{}
		err := json.NewDecoder(r.Body).Decode(&userChecking)
		logFatal(err)
		userRepo := userRepository.UserRepository{}
		userFromBase.Password = userRepo.Signin(db, userChecking, userFromBase)
		userOK := true
		if err = bcrypt.CompareHashAndPassword([]byte(userFromBase.Password), []byte(userChecking.Password)); err != nil {
			userOK = false
			json.NewEncoder(w).Encode(&userOK)
		}
		json.NewEncoder(w).Encode(&userOK)
	}
}
