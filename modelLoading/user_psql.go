package modelLoading

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"knock-knock/models"
	//"log"
)

func LoadUsers(db *sqlx.DB) {
	rows, err := db.Query("select id_user, firstname, lastname, summarymark, teamcount, CASE WHEN teamcount>0 THEN (summarymark::decimal /teamcount) ELSE 0 END AS avaragemark from teammate")
	logErr(err)
	defer rows.Close()
	users := []*models.User{}
	err = sqlx.StructScan(rows, &users)
	logErr(err)
	for _, u := range users {
		models.AddUser(u)
	}
}

func Signup(db *sqlx.DB, user models.User) int {
	err := db.QueryRow("insert into teammate (firstname,lastname,password) values ($1,$2,$3) RETURNING id_user;",
		user.Firstname, user.Lastname, user.Password).Scan(&user.ID)
	logErr(err)
	models.AddUser(&user)
	return user.ID
}

func Signin(db *sqlx.DB, userСhecking models.User, userFromBase models.User) (string, error) {
	err := db.QueryRow("select password from teammate where id_user=$1", userСhecking.ID).Scan(&userFromBase.Password)
	if err == sql.ErrNoRows {
		return "", err
	}
	logErr(err)
	return userFromBase.Password, nil
}
