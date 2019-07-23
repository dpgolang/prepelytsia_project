package models

import "sort"

type User struct {
	ID          int     `json:"id" db:"id_user"`
	Firstname   string  `json:"firstname,omitempty" db:"firstname"`
	Lastname    string  `json:"lastname,omitempty" db:"lastname"`
	Sum         int     `json:"sum,omitempty" db:"summarymark"`
	Teamcount   int     `json:"teamcount,omitempty"db:"teamcount"`
	AvarageMark float32 `json:"avarageMark,omitempty" db:"avaragemark"`
	Password    string  `json:"password,omitempty" db:"password"`
	Teams       []*Team
}

type UserShow struct {
	ID          int     `json:"id"`
	Firstname   string  `json:"firstname,omitempty"`
	Lastname    string  `json:"lastname,omitempty`
	Sum         int     `json:"sum,omitempty"`
	Teamcount   int     `json:"teamcount,omitempty"`
	AvarageMark float32 `json:"avarageMark,omitempty"`
	TeamIDs     []int   `json:"teamIds"`
}

func (user *User) ToUserShow() UserShow {
	userShow := UserShow{
		ID:          user.ID,
		Firstname:   user.Firstname,
		Lastname:    user.Lastname,
		Sum:         user.Sum,
		Teamcount:   user.Teamcount,
		AvarageMark: user.AvarageMark,
	}
	for _, t := range user.Teams {
		userShow.TeamIDs = append(userShow.TeamIDs, t.ID)
	}
	return userShow
}

var Users = make(map[int]*User)

func GetUsers() []*User {
	users := []*User{}
	for _, u := range Users {
		users = append(users, u)
	}
	sort.Slice(users, func(i, j int) bool {
		return users[i].ID > users[j].ID
	})
	return users
}

func GetUser(id int) *User {
	return Users[id]
}

func AddUser(u *User) {
	Users[u.ID] = u
}

func (u *User) CalculateAvarageMark() {
	if u.Teamcount == 0 {
		u.AvarageMark = float32(0)
		return
	}
	u.AvarageMark = float32(u.Sum) / float32(u.Teamcount)
}
