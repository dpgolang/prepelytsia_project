package controllers

import (
	"fmt"
	_ "database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	_ "html/template"
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

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Print secret message
	///fmt.Fprintln(w, "The cake is a lie!")
}
func (c Controller) GetUsers(db *sqlx.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "cookie-name")

		// Check if user is authenticated
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		//var user models.User
		userRepo := userRepository.UserRepository{}
		users, err := userRepo.GetUsers(db)
		logFatal(err)
		/*data := struct {
			Title   string
			Content []models.User
		}{
			Title:   "Users",
			Content: users,
		}
		tmpl, _ := template.ParseFiles("templates/users.html")
		tmpl.Execute(w, data)*/
		json.NewEncoder(w).Encode(users)
	}
}

func (c Controller) GetUser(db *sqlx.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "cookie-name")

		// Check if user is authenticated
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		var user models.User
		userRepo := userRepository.UserRepository{}
		id, err := strconv.Atoi(params["id"])
		logFatal(err)
		user, userIsFound := userRepo.GetUser(db, user, id)
		if userIsFound {
			json.NewEncoder(w).Encode(user)
			return
		}

		w.WriteHeader(http.StatusUnauthorized)
		/*if user.Id == 0 {
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
		tmpl.Execute(w, data)*/
	}
}

func (c Controller) Signup(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		var userID int
		w.Header().Set("Content-Type", "application/json")
		err := json.NewDecoder(r.Body).Decode(&user)
		logFatal(err)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
		user.Password = string(hashedPassword)
		userRepo := userRepository.UserRepository{}
		userID = userRepo.Signup(db, user)
		json.NewEncoder(w).Encode(&userID)
	}
}
func (c Controller) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("aasdadsa")
		session, _ := store.Get(r, "cookie-name")
		
		session.Options.MaxAge = -1
		err := session.Save(r, w)
		if err != nil {
			logFatal(err)
		}
		json.NewEncoder(w).Encode("logged out!")
	}
}
func (c Controller) Signin(db *sqlx.DB) http.HandlerFunc { //как отлавливать когда айди не найдет
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "cookie-name")
		session.Options.MaxAge = 300
		w.Header().Set("Content-Type", "application/json")
		userFromBase := models.User{}
		userChecking := models.User{}
		var userIsFound bool
		err := json.NewDecoder(r.Body).Decode(&userChecking)
		logFatal(err)
		userRepo := userRepository.UserRepository{}
		userFromBase.Password, userIsFound = userRepo.Signin(db, userChecking, userFromBase)
		if userIsFound {
			if err = bcrypt.CompareHashAndPassword([]byte(userFromBase.Password), []byte(userChecking.Password)); err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				//json.NewEncoder(w).Encode()
				return
			}
			session.Values["authenticated"] = true
			session.Save(r, w)
			json.NewEncoder(w).Encode(fmt.Sprintf("logged in, id: %d",userChecking.Id))
			//json.NewEncoder(w).Encode(userChecking.Id)
		}
	}
}
