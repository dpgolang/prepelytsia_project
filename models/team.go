package models

import (
	"sort"
)

type Team struct {
	ID        int    `json:"id" db:"id_team"`
	Name      string `json:"name" db:"name"`
	IDcreator int    `json:"creator" db:"id_creator"`
	Done      bool   `json:"done" db:"done"`
	Users     []*User
	UsersDone map[int]bool
	Marks     map[int][]Mark
}

type TeamShow struct {
	ID        int
	Name      string
	CreatorId int
	Done      bool
	Users     []UserShow
	Marks     map[int][]Mark
}

type Mark struct {
	UserID int
	Value  uint8
}

var Teams = make(map[int]*Team)

func (t *Team) ToTeamShow() TeamShow {
	teamShow := TeamShow{
		ID:        t.ID,
		Done:      t.IsDone(),
		Name:      t.Name,
		CreatorId: t.IDcreator,
		Marks:     t.Marks,
	}
	for _, u := range t.Users {
		teamShow.Users = append(teamShow.Users, u.ToUserShow())
	}
	return teamShow
}

func GetTeam(id int) *Team {
	return Teams[id]
}

func GetTeams() []*Team {
	teams := []*Team{}
	for _, t := range Teams {
		teams = append(teams, t)
	}
	return SortTeams(teams)
}

func GetDoneTeams(done bool) []*Team {
	teams := []*Team{}
	for _, t := range Teams {
		if (done && t.IsDone()) || (!done && !t.IsDone()) {
			teams = append(teams, t)
		}
	}
	return SortTeams(teams)
}

func SortTeams(teams []*Team) []*Team {
	sort.Slice(teams, func(i, j int) bool {
		return teams[i].ID > teams[j].ID
	})
	return teams
}

func AddTeam(t *Team) {
	Teams[t.ID] = t
}

func (t *Team) IsDone() bool {
	if len(t.Users) == 0 {
		return false
	}
	for _, u := range t.Users {
		if !t.UsersDone[u.ID] {
			return false
		}
	}
	return true
}

func (t *Team) isUserPresent(u *User) bool {
	for _, user := range t.Users {
		if user.ID == u.ID {
			return true
		}
	}
	return false
}

func (t *Team) AddUser(u *User) {
	if t.isUserPresent(u) {
		return
	}
	t.Users = append(t.Users, u)
	u.Teams = append(u.Teams, t)
}

func (t *Team) SetUserDone(u *User) {
	if t.isUserPresent(u) {
		t.UsersDone[u.ID] = true
	}
}

func (t *Team) SetMark(fromUser, toUser *User, mark uint8) bool {
	if !t.isUserPresent(fromUser) || !t.isUserPresent(toUser) {
		return false
	}

	if t.Marks == nil {
		t.Marks = make(map[int][]Mark)
	}

	marks := t.Marks[fromUser.ID]
	for _, mark := range marks {
		if mark.UserID == toUser.ID {
			return false
		}
	}

	newMark := Mark{
		UserID: toUser.ID,
		Value:  mark,
	}

	marks = append(marks, newMark)
	t.Marks[fromUser.ID] = marks
	return true
}
