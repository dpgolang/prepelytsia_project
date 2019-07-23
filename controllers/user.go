package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/subosito/gotenv"
	"golang.org/x/crypto/bcrypt"
	_ "html/template"
	"knock-knock/modelLoading"
	"knock-knock/models"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func init() {
	gotenv.Load()
	store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
}

//var users []models.User

type Controller struct{}

func logErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

//куда лучше вставить эту строчку? как подгрузить из env
var store *sessions.CookieStore

func (c Controller) GetMyPage(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "cookie-name")
		w.Header().Set("Content-Type", "application/json")
		userCheckingID := getUserSession(session)
		if userCheckingID != 0 {
			json.NewEncoder(w).Encode(userCheckingID)
			user := models.GetUser(userCheckingID)
			if user != nil {
				json.NewEncoder(w).Encode(user.ToUserShow())
				return
			}
			json.NewEncoder(w).Encode("There is no such user!")
			http.Error(w, "You should sign in to check this page", http.StatusForbidden)
		}
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
		users := models.GetUsers()
		usersShow := []models.UserShow{}
		for _, user := range users {
			usersShow = append(usersShow, user.ToUserShow())
		}
		/*data := struct {
			Title   string
			Content []models.User
		}{
			Title:   "Users",
			Content: users,
		}
		tmpl, _ := template.ParseFiles("templates/users.html")
		tmpl.Execute(w, data)*/
		json.NewEncoder(w).Encode(usersShow)
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

		id, err := strconv.Atoi(params["id"])
		logErr(err)
		user := models.GetUser(id)
		if user != nil {
			json.NewEncoder(w).Encode(user.ToUserShow())
			return
		}
		json.NewEncoder(w).Encode("There is no such user!")
		w.WriteHeader(http.StatusNotFound)

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
		logErr(err)
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
		logErr(err)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if checkName(user) {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
			logErr(err)
			user.Password = string(hashedPassword)
			userID = modelLoading.Signup(db, user)
			json.NewEncoder(w).Encode(&userID)
			return
		}
		http.Error(w, "Wrong firstname or lastname!", http.StatusForbidden)
	}
}

func checkName(user models.User) bool {
	matchedFirstname, err := regexp.MatchString(`^[a-zA-Z0-9_а-яА-Я ]+$`, user.Firstname)
	logErr(err)
	matchedLastname, err := regexp.MatchString(`^[a-zA-Z0-9_а-яА-Я ]+$`, user.Lastname)
	return matchedLastname && matchedFirstname
}

func (c Controller) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "cookie-name")
		session.Options.MaxAge = -1
		err := session.Save(r, w)
		if err != nil {
			logErr(err)
		}

		json.NewEncoder(w).Encode("logged out!")
	}
}
func (c Controller) Signin(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "cookie-name")
		session.Options.MaxAge = 300
		w.Header().Set("Content-Type", "application/json")
		userFromBase := models.User{}
		userChecking := models.User{}
		err := json.NewDecoder(r.Body).Decode(&userChecking)
		logErr(err)
		userFromBase.Password, err = modelLoading.Signin(db, userChecking, userFromBase)
		if err == nil {
			if err = bcrypt.CompareHashAndPassword([]byte(userFromBase.Password), []byte(userChecking.Password)); err != nil {
				http.Error(w, "Wrong ID or password!", http.StatusUnauthorized)
				return
			}
			session.Values["authenticated"] = true
			session.Values["id"] = userChecking.ID
			session.Save(r, w)
			json.NewEncoder(w).Encode(fmt.Sprintf("logged in, id: %d", userChecking.ID))
		}
		if err == sql.ErrNoRows {
			http.Error(w, "Wrong ID or password!", http.StatusUnauthorized)
			return
		}
		logErr(err)
	}
}
