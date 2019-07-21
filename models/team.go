package models

type Team struct {
	Id         int    `json:"id" db:"id_team"`
	Name       string `json:"name" db:"name"`
	Id_creator int    `json:"creator" db:"id_creator"`
	Done       bool   `json:"done" db:"done"`
}
