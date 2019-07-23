package models

type Teams map[int]map[int]User

type TeamList struct {
	//TeamNumber int    `json:"teamnumber" db:"team_id"`
	Teammates []User //`json: "teammates" db: "teammates"`
	TeamName  string //`json: teamname db:"teamname"`
	Done      bool   //`json: done db: "done"`
	CreatorID int    //`json: creatorid db: "id_creator"`
	/*IDelem     int    `json:"id" db:"id_elem"`
	  IDteam     string `json:"name" db:"id_team"`
	  IDteammate int    `json:"creator" db:"id_teammate"`
	  Done       bool   `json:"done" db:"done"`*/
}
