package models

type TeamList struct {
	IDelem     int    `json:"id" db:"id_elem"`
	IDteam     string `json:"name" db:"id_team"`
	IDteammate int    `json:"creator" db:"id_teammate"`
	Done       bool   `json:"done" db:"done"`
}
