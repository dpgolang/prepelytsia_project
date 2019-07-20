package userRepository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"knock-knock/models"
	"log"
)

type UserRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (u UserRepository) GetUsers(db *sqlx.DB) ([]models.User, error) { //ддобавить сканер интерфейс для юзера чтобы считать среднее арифметическое
	rows, err := db.Query("select id_user, firstname, lastname, summarymark, teamcount  from teammate")
	users := []models.User{}
	logFatal(err)
	defer rows.Close()
	for rows.Next() {
		err := sqlx.StructScan(rows, &users)
		if err != nil {
			return []models.User{}, err
		}
	}
	//users.CalculateAvarageMark()
	/*for i := range users {
		users[i].CalculateAvarageMark()
	}*/
	return users, nil
}

func (u UserRepository) GetUser(db *sqlx.DB, user models.User, id int) (models.User, bool) {
	row := db.QueryRowx("select id_user, firstname, lastname, summarymark, teamcount from teammate where id_user = $1", id)
	err := row.StructScan(&user)
	if err == sql.ErrNoRows {
		return models.User{}, false
	}
	logFatal(err)
	user.CalculateAvarageMark()
	return user, true
}

func (u UserRepository) Signup(db *sqlx.DB, user models.User) int {
	err := db.QueryRow("insert into teammate (firstname,lastname,password) values ($1,$2,$3) RETURNING id_user;",
		user.Firstname, user.Lastname, user.Password).Scan(&user.Id)
	logFatal(err)
	return user.Id
}
func (u UserRepository) Signin(db *sqlx.DB, userСhecking models.User, userFromBase models.User) (string, bool) {
	err := db.QueryRow("select password from teammate where id_user=$1", userСhecking.Id).Scan(&userFromBase.Password) //вернуть
	if err == sql.ErrNoRows {
		return "", false
	}
	logFatal(err)
	return userFromBase.Password, true
}
