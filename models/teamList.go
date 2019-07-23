package models

//type Teams map[int]map[int]User

type TeamMember struct {
	Team     Team `json:"team,omitempty"`
	Teammate User `json:"teammate,omitempty" db:"teammate"`
	UserDone bool `json:"userdone"db:"userdone"`
}
type TeamListUser struct {
	Teammate User `json:"teammate,omitempty" db:"teammate"`
	UserDone bool `json:"userdone"db:"userdone"`
}
type ParsedMember struct {
	TeamFromList Team           `json:"team"`
	Teammates    []TeamListUser `json:"teammates"`
}
