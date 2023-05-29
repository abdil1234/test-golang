package models

type Game struct {
	Id      int    `db:"id" json:"id"`
	Mdate   int64  `db:"mdate" json:"mdate"`
	Stadium string `db:"stadium" json:"stadium"`
	Team1   string `db:"team1" json:"team1"`
	Team2   string `db:"team2" json:"team2"`
}

func (b Game) TableName() string {
	return "game"
}
