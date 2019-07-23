package teamListRepository

import (
	_ "database/sql"
	"github.com/jmoiron/sqlx"
	"knock-knock/models"
	"log"
)

type TeamListRepository struct{}

func logErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

// func (u TeamRepository) GetAllTeamList(db *sqlx.DB) ([]TeamList, error) {
// 	rows, err := db.Query("select id_team, from team")
// 	teams := []models.Team{}
// 	logErr(err)
// 	defer rows.Close()
// 	err = sqlx.StructScan(rows, &teams)
// 	if err != nil {
// 		return []models.Team{}, err
// 	}
// 	return teams, nil
// }

func (u TeamListRepository) GetTeamList(db *sqlx.DB) ([]models.TeamList, error) {
	var teams []models.TeamList
	sql := `SELECT
      (teammate.*) as teammates,
      (team.name) as teamname,
      (team.id_creator) as id_creator,
      (team.done) as done
    FROM
      teamlist
    INNER JOIN teammate ON teammate.id_user = teamlist.id_teammate
    INNER JOIN team ON teamlist.id_team = team.id_team`
	err := db.Select(&teams, sql)
	return teams, err
}

// func (u TeamRepository) GetTeam(db *sqlx.DB, team models.Team, id int) (models.Team, error) {
// 	row := db.QueryRowx("select * from team where id_team = $1", id)
// 	err := row.StructScan(&team)
// 	if err == sql.ErrNoRows {
// 		return models.Team{}, err
// 	}
// 	logErr(err)
// 	return team, nil
// }

// func (u TeamRepository) GetMyTeams(db *sqlx.DB, team models.Team, id_user int, id_creator int) (models.Team, bool) {
//  	row := db.QueryRowx("select * from team where id_user = $1", id)
//  	err := row.StructScan(&team)
//  	if err == sql.ErrNoRows {
//  		return models.Team{}, false
//  	}
//  	logErr(err)
//  	return team, true
// }
