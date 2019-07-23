package modelLoading

import (
	//"database/sql"
	"github.com/jmoiron/sqlx"
	"knock-knock/models"
	//"log"
)

func LoadTeams(db *sqlx.DB) {
	rows, err := db.Query("select * from team")
	teams := []*models.Team{}
	logErr(err)
	defer rows.Close()
	err = sqlx.StructScan(rows, &teams)
	logErr(err)
	for _, t := range teams {
		models.AddTeam(t)
	}

	rowsTeamlist, err := db.Query(`SELECT id_team, id_teammate, done FROM teamlist`)
	logErr(err)
	defer rowsTeamlist.Close()
	for rowsTeamlist.Next() {
		var (
			teamId int
			userId int
			done   bool
		)
		err := rowsTeamlist.Scan(&teamId, &userId, &done)
		logErr(err)
		user := models.GetUser(userId) //отбросили невалидную инфу
		if user == nil {
			continue
		}
		team := models.GetTeam(teamId)
		if team == nil {
			continue
		}
		team.AddUser(user)
		if done {
			team.SetUserDone(user)
		}
	}
}

// func (u TeamRepository) GetMyTeams(db *sqlx.DB, team models.Team, id_user int, id_creator int) (models.Team, bool) {
//  	row := db.QueryRowx("select * from team where id_user = $1", id)
//  	err := row.StructScan(&team)
//  	if err == sql.ErrNoRows {
//  		return models.Team{}, false
//  	}
//  	logErr(err)
//  	return team, true
// }
