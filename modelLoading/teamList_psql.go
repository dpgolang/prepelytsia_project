package modelLoading

import (
	_ "database/sql"
	_ "github.com/jmoiron/sqlx"
	_ "knock-knock/models"
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

// func (u TeamListRepository) GetTeamList(db *sqlx.DB) ([]models.Team, error) {
// 	rows, err := db.Query(`SELECT team.id_team, team.name, team.done, team.Id_creator,
// 			teammate.id_user, teammate.firstname, teammate.lastname, teamlist.done
// 		FROM teamlist
// 		JOIN teammate ON teammate.id_user = teamlist.id_teammate
// 		JOIN team ON teamlist.id_team = team.id_team
// 		ORDER BY teamlist.id_team`)
// 	logErr(err)
// 	defer rows.Close()
// 	teamList := []models.Team{}
// 	for rows.Next() {
// 		member := models.TeamMember{}
// 		err = rows.Scan(&m.Team.IDteam, &m.Team.Name, &m.Team.Done, &m.Team.IDcreator,
// 			&m.Teammate.IDuser, &m.Teammate.Firstname, &m.Teammate.Lastname, &m.UserDone)
// 		logErr(err)
// 		teamList = append(teamList, m)
// 	}
// 	return teamList, err
// }

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
