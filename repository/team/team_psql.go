package teamRepository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"knock-knock/models"
	"log"
)

type TeamRepository struct{}

func logErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func (u TeamRepository) GetTeams(db *sqlx.DB) ([]models.Team, error) {
	rows, err := db.Query("select * from team")
	teams := []models.Team{}
	logErr(err)
	defer rows.Close()
	err = sqlx.StructScan(rows, &teams)
	if err != nil {
		return []models.Team{}, err
	}
	return teams, nil
}

func (u TeamRepository) GetTeamsDone(db *sqlx.DB, done bool) ([]models.Team, error) {
	rows, err := db.Query("select * from team where done = $1", done)
	teams := []models.Team{}
	logErr(err)
	defer rows.Close()
	err = sqlx.StructScan(rows, &teams)
	if err != nil {
		return []models.Team{}, err
	}
	return teams, nil
}

func (u TeamRepository) GetTeam(db *sqlx.DB, team models.Team, id int) (models.Team, error) {
	row := db.QueryRowx("select * from team where id_team = $1", id)
	err := row.StructScan(&team)
	if err == sql.ErrNoRows {
		return models.Team{}, err
	}
	logErr(err)
	return team, nil
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
