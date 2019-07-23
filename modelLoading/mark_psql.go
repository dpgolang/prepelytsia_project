package modelLoading

import (
	//"database/sql"
	"github.com/jmoiron/sqlx"
	"knock-knock/models"
	//"log"
)

func LoadMarks(db *sqlx.DB) {
	rowsMarklist, err := db.Query(`SELECT id_team, id_user_get_mark, id_user_give_mark, mark FROM marks`)
	logErr(err)
	defer rowsMarklist.Close()
	for rowsMarklist.Next() {
		var (
			teamID         int
			userGetMarkID  int
			userGiveMarkID int
			mark           uint8
		)
		err = rowsMarklist.Scan(&teamID, &userGetMarkID, &userGiveMarkID, &mark)
		logErr(err)
		userGetMark := models.GetUser(userGetMarkID) //отбросили невалидную инфу
		if userGetMark == nil {
			continue
		}
		userGiveMark := models.GetUser(userGiveMarkID)
		if userGiveMark == nil {
			continue
		}
		team := models.GetTeam(teamID)
		if team == nil {
			continue
		}
		team.SetMark(userGiveMark, userGetMark, mark)
	}
}
