package models

type User struct {
	IDuser      int     `json:"id" db:"id_user"`
	Firstname   string  `json:"firstname,omitempty" db:"firstname"`
	Lastname    string  `json:"lastname,omitempty" db:"lastname"`
	Sum         int     `json:"sum" db:"summarymark"`
	Teamcount   int     `json:"teamcount"db:"teamcount"`
	AvarageMark float32 `json:"avarageMark"` //`db:"telcode"`
	Password    string  `json:"password,omitempty" db:"password"`
}

type Users []*User

func (u Users) CalculateAvarageMark() { // считать в бд
	for _, i := range u {
		i.CalculateAvarageMark()
	}
}

func (u *User) CalculateAvarageMark() {
	if u.Teamcount == 0 {
		u.AvarageMark = float32(0)
		return
	}
	u.AvarageMark = float32(u.Sum) / float32(u.Teamcount)
}
