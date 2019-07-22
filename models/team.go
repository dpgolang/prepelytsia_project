package models

type Team struct {
	IDteam    int    `json:"id" db:"id_team"`
	Name      string `json:"name" db:"name"`
	IDcreator int    `json:"creator" db:"id_creator"`
	Done      bool   `json:"done" db:"done"`
}
