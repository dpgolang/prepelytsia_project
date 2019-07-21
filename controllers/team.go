package controllers

import (
	_ "database/sql"
	"encoding/json"
	_ "fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	_ "github.com/subosito/gotenv"
	_ "golang.org/x/crypto/bcrypt"
	_ "html/template"
	"knock-knock/models"
	"knock-knock/repository/team"
	"log"
	"net/http"
	//"os"
	_ "regexp"
	"strconv"
)

// func (c Controller) GetMyTeams(db *sqlx.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		session, _ := store.Get(r, "cookie-name")
// 		session.Options.MaxAge = 300
// 		w.Header().Set("Content-Type", "application/json")
// 		userCheckingID := getUserSession(session)
// 		var userIsFound bool
// 		var team models.Team
// 		//fmt.Println("aaa")

// 		//fmt.Println(userCheckingID.(int))
// 		userRepo := userRepository.UserRepository{}
// 		if userCheckingID != 0 {
// 			json.NewEncoder(w).Encode(userCheckingID)
// 			team, userIsFound = userRepo.GetUser(db, team, userCheckingID)
// 			fmt.Println(team)
// 			if userIsFound {
// 				json.NewEncoder(w).Encode(team)
// 				return
// 			}
// 		}
// 		//	session.Values["authenticated"] = true
// 		http.Error(w, "You should sign in to check this page", http.StatusForbidden)
// 		//json.NewEncoder(w).Encode(fmt.Sprintf("logged in, id: %d",userChecking.Id))
// 		//json.NewEncoder(w).Encode(userChecking.Id)
// 	}
// }

// func getUserSession(s *sessions.Session) int {
// 	userID, ok := s.Values["id"].(int)
// 	if !ok {
// 		return 0
// 	}
// 	return userID
// }

func (c Controller) GetTeams(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		session, _ := store.Get(r, "cookie-name")
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Error(w, "You should sign in to check this page", http.StatusForbidden)
			return
		}
		showAllDone, _ := strconv.ParseBool(r.URL.Query().Get("showAllDone"))
		showAllNotDone, _ := strconv.ParseBool(r.URL.Query().Get("showAllNotDone"))
		var (
			err      error
			teamRepo teamRepository.TeamRepository
			teams    []models.Team
		)
		if showAllDone {
			teams, err = teamRepo.GetTeamsDone(db, showAllDone)
		} else if showAllNotDone {
			teams, err = teamRepo.GetTeamsDone(db, !showAllNotDone)
		} else {
			teams, err = teamRepo.GetTeams(db)
		}
		logFatal(err)
		json.NewEncoder(w).Encode(teams)
		log.Println(teams)
	}
}

func (c Controller) GetTeam(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "cookie-name")
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Error(w, "You should sign in to check this page", http.StatusForbidden)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		var team models.Team
		teamRepo := teamRepository.TeamRepository{}
		id, err := strconv.Atoi(params["id"])
		logFatal(err)
		team, teamIsFound := teamRepo.GetTeam(db, team, id)
		if teamIsFound {
			json.NewEncoder(w).Encode(team)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
	}
}
