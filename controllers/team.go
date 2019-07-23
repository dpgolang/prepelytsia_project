package controllers

import (
	//"database/sql"
	"encoding/json"
	_ "fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	_ "github.com/subosito/gotenv"
	_ "golang.org/x/crypto/bcrypt"
	_ "html/template"
	"knock-knock/models"
	//"knock-knock/repository/team"
	"log"
	"net/http"
	//"os"
	_ "regexp"
	"strconv"
)

func (c Controller) GetTeams(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		session, _ := store.Get(r, "cookie-name")
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Error(w, "You should sign in to check this page", http.StatusForbidden)
			return
		}
		var (
			err   error
			teams []*models.Team
		)
		done, err := strconv.ParseBool(r.URL.Query().Get("done"))
		if err != nil {
			teams = models.GetTeams()
		} else {
			teams = models.GetDoneTeams(done)
		}

		teamsShow := []models.TeamShow{}
		for _, t := range teams {
			teamsShow = append(teamsShow, t.ToTeamShow())
		}

		logErr(err)
		json.NewEncoder(w).Encode(teamsShow)
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
		id, err := strconv.Atoi(params["id"])
		logErr(err)
		team := models.GetTeam(id)
		if team != nil {
			json.NewEncoder(w).Encode(team.ToTeamShow())
			return
		}
		json.NewEncoder(w).Encode("There is no such team!")
		w.WriteHeader(http.StatusNotFound)
	}
}
