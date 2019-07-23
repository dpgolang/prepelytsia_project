package controllers

import (
	_ "database/sql"
	"encoding/json"
	_ "fmt"
	_ "github.com/gorilla/mux"
	_ "github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	_ "github.com/subosito/gotenv"
	_ "golang.org/x/crypto/bcrypt"
	_ "html/template"
	"knock-knock/models"
	"knock-knock/repository/teamList"
	"log"
	"net/http"
	//"os"
	_ "regexp"
	_ "strconv"
)

func (c Controller) GetTeamList(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*	w.Header().Set("Content-Type", "application/json")
			session, _ := store.Get(r, "cookie-name")
			if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
				http.Error(w, "You should sign in to check this page", http.StatusForbidden)
				return
			}*/
		var (
			err          error
			teamListRepo teamListRepository.TeamListRepository
			teamList     []models.TeamList
		)
		teamList, err = teamListRepo.GetTeamList(db)
		//teams, err = teamRepo.GetTeamsDone(db, done)
		logErr(err)
		json.NewEncoder(w).Encode(teamList)
		log.Println(teamList)
	}
}
