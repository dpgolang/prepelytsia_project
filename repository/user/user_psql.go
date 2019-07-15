package userRepository

import (
	"database/sql"
	"knock-knock/models"
	"log"
)

type UserRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (u UserRepository) GetUsers(db *sql.DB, user models.User, users []models.User) []models.User {
	rows, err := db.Query("select * from teammate")
	logFatal(err)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Sum, &user.Teamcount, &user.Password)
		logFatal(err)
		user.CalculateAvarageMark()
		users = append(users, user)
	}
	return users
}

func (u UserRepository) GetUser(db *sql.DB, user models.User, id int) models.User {
	row := db.QueryRow("select * from teammate where id_user = $1", id)
	err := row.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Sum, &user.Teamcount, &user.Password)
	if user.Id == 0 {
		return models.User{}
	}
	logFatal(err)
	user.CalculateAvarageMark()
	return user
}

func (u UserRepository) Signup(db *sql.DB, user models.User) int {
	err := db.QueryRow("insert into teammate (firstname,lastname,password) values ($1,$2,$3) RETURNING id_user;",
		user.Firstname, user.Lastname, user.Password).Scan(&user.Id)
	logFatal(err)
	return user.Id
}
func (u UserRepository) Signin(db *sql.DB, userСhecking models.User, userFromBase models.User) string {
	err := db.QueryRow("select password from teammate where id_user=$1", userСhecking.Id).Scan(&userFromBase.Password) //как отлавливать когда айди не найден
	logFatal(err)
	return userFromBase.Password
}
