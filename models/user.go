package models

type User struct {
	Id          int     `json:"id" db:"id_user"`
	Firstname   string  `json:"firstname,omitempty"`
	Lastname    string  `json:"lastname,omitempty"`
	Sum         int     `json:"sum,omitempty" db:"summarymark"`
	Teamcount   int     `json:"teamcount,omitempty"`
	AvarageMark float32 `json:"avarageMark,omitempty"` //`db:"telcode"`
	Password string `json:"password,omitempty"`
}



type Users []*User

func(u Users) CalculateAvarageMark(){ // почему не работает в repository/user/user_psql в getUsers users.CalculateAvarageMark()?
	for _, i := range u {
		i.CalculateAvarageMark()
	}
}

func (u *User) CalculateAvarageMark() {
	if u.Teamcount == 0 {
		u.AvarageMark = float32(1)
	}
	u.AvarageMark = float32(u.Sum) / float32(u.Teamcount)
}

