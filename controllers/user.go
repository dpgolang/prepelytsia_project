package controllers

import (
	_ "database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	_ "github.com/subosito/gotenv"
	"golang.org/x/crypto/bcrypt"
	_ "html/template"
	"knock-knock/models"
	"knock-knock/repository/user"
	"log"
	"net/http"
	//"os"
	"regexp"
	"strconv"
)

/*func init() {
	gotenv.Load()
}*/

//var users []models.User

type Controller struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//куда лучше вставить эту строчку? как подгрузить из env
var store = sessions.NewCookieStore([]byte("secret-key"))

func (c Controller) GetMyPage(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "cookie-name")
		w.Header().Set("Content-Type", "application/json")
		userCheckingID := getUserSession(session)
		var (
			userIsFound bool
			user        models.User
		)
		userRepo := userRepository.UserRepository{}
		if userCheckingID != 0 {
			json.NewEncoder(w).Encode(userCheckingID)
			user, userIsFound = userRepo.GetUser(db, user, userCheckingID)
			fmt.Println(user)
			if userIsFound {
				json.NewEncoder(w).Encode(user)
				return
			}
		}
		//	session.Values["authenticated"] = true
		http.Error(w, "You should sign in to check this page", http.StatusForbidden)
		//json.NewEncoder(w).Encode(fmt.Sprintf("logged in, id: %d",userChecking.Id))
		//json.NewEncoder(w).Encode(userChecking.Id)
	}
}

func getUserSession(s *sessions.Session) int {
	userID, ok := s.Values["id"].(int)
	if !ok {
		return 0
	}
	return userID
}

func (c Controller) GetUsers(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "cookie-name")
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Error(w, "You should sign in to check this page", http.StatusForbidden)
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
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Error(w, "You should sign in to check this page", http.StatusForbidden)
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
		var (
			user   models.User
			userID int
		)
		w.Header().Set("Content-Type", "application/json")
		err := json.NewDecoder(r.Body).Decode(&user)
		logFatal(err)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		nameOK := checkName(user)
		if nameOK {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
			logFatal(err)
			user.Password = string(hashedPassword)
			userRepo := userRepository.UserRepository{}
			userID = userRepo.Signup(db, user)
			json.NewEncoder(w).Encode(&userID)
			return
		}
		http.Error(w, "Wrong firstname or lastname!", http.StatusForbidden)
	}
}

func checkName(user models.User) bool {
	matchedFirstname, err := regexp.MatchString(`^[a-zA-Z0-9_а-яА-Я \-]*$`, user.Firstname)
	logFatal(err)
	matchedLastname, err := regexp.MatchString(`^[a-zA-Z0-9_а-яА-Я \-]*$`, user.Lastname)
	return matchedLastname && matchedFirstname
}

func (c Controller) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
				http.Error(w, "Wrong password!", http.StatusUnauthorized)
				return
			}
			session.Values["authenticated"] = true
			session.Values["id"] = userChecking.Id
			session.Save(r, w)
			json.NewEncoder(w).Encode(fmt.Sprintf("logged in, id: %d", userChecking.Id))
		} else {
			http.Error(w, "Wrong ID!", http.StatusUnauthorized)
			return
		}
	}
}
