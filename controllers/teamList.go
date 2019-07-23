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
	"net/http"
	//"os"
	"log"
	_ "regexp"
	_ "strconv"
)

func (c Controller) GetTeamList(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		/*	w.Header().Set("Content-Type", "application/json")
			session, _ := store.Get(r, "cookie-name")
			if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
				http.Error(w, "You should sign in to check this page", http.StatusForbidden)
				return
			}*/
		var (
			err          error
			teamListRepo teamListRepository.TeamListRepository
			listMember   []models.TeamMember
			parsedList   []models.ParsedMember
		)
		listMember, err = teamListRepo.GetTeamList(db)
		//teams, err = teamRepo.GetTeamsDone(db, done)
		logErr(err)
		parsedList = parseList(listMember)
		json.NewEncoder(w).Encode(parsedList)
		//log.Println(parsedList)
	}
}

func parseList(list []models.TeamMember) []models.ParsedMember {
	var (
		parsedList []models.ParsedMember
		//parsedListUnit models.ParsedMember
		user  models.TeamListUser
		users []models.TeamListUser
		//team           []models.Team
	)
	for i := 0; i < len(list); i++ {
		user.Teammate = list[i].Teammate
		user.UserDone = list[i].UserDone
		users = append(users, user)
		if i+1 < len(list) && list[i].Team.IDteam != list[i+1].Team.IDteam {
			parsedListUnit := models.ParsedMember{}
			parsedListUnit.TeamFromList = list[i].Team
			parsedListUnit.Teammates = users
			parsedList = append(parsedList, parsedListUnit)
			users = []models.TeamListUser{}
		}
	}
	log.Println(len(list))
	return parsedList
}
