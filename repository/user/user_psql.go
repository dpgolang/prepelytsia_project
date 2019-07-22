package userRepository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"knock-knock/models"
	"log"
)

type UserRepository struct{}

func logErr(err error) {
	if err != nil {
		log.Print(err)
	}
}

func (u UserRepository) GetUsers(db *sqlx.DB) ([]models.User, error) {
	rows, err := db.Query("select id_user, firstname, lastname, summarymark, teamcount  from teammate")
	users := []models.User{}
	logErr(err)
	defer rows.Close()
	err = sqlx.StructScan(rows, &users)
	if err != nil {
		return []models.User{}, err
	}
	//users.CalculateAvarageMark()
	for i := range users {
		users[i].CalculateAvarageMark()
	}
	return users, nil
}

func (u UserRepository) GetUser(db *sqlx.DB, user models.User, id int) (models.User, error) {
	row := db.QueryRowx("select id_user, firstname, lastname, summarymark, teamcount from teammate where id_user = $1", id)
	err := row.StructScan(&user)
	if err == sql.ErrNoRows {
		return models.User{}, err
	}
	logErr(err)
	user.CalculateAvarageMark()
	return user, err
}

func (u UserRepository) Signup(db *sqlx.DB, user models.User) int {
	err := db.QueryRow("insert into teammate (firstname,lastname,password) values ($1,$2,$3) RETURNING id_user;",
		user.Firstname, user.Lastname, user.Password).Scan(&user.IDuser)
	logErr(err)
	return user.IDuser
}
func (u UserRepository) Signin(db *sqlx.DB, userСhecking models.User, userFromBase models.User) (string, error) {
	err := db.QueryRow("select password from teammate where id_user=$1", userСhecking.IDuser).Scan(&userFromBase.Password)
	if err == sql.ErrNoRows {
		return "", err
	}
	logErr(err)
	return userFromBase.Password, nil
}
