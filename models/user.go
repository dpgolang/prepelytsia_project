package models

type User struct {
	Id          int     `json:"id"`
	Firstname   string  `json:"firstname"`
	Lastname    string  `json:"lastname"`
	Sum         int     `json:"sum"`
	Teamcount   int     `json:"teamcount"`
	AvarageMark float32 `json:"avarageMark"`
	Password string `json:"password"`
}

func (u *User) CalculateAvarageMark() {
	u.AvarageMark = float32(u.Sum) / float32(u.Teamcount)
}