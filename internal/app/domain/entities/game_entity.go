package entities

import (
	"time"

	"github.com/abdil1234/test-golang/internal/app/domain/repositories/models"
)

type GameModel struct {
	Game models.Game
}
type GameEntity struct {
	Id      int    `json:"id"`
	Mdate   string `json:"mdate"`
	Stadium string `json:"stadium"`
	Team1   string `json:"team1"`
	Team2   string `json:"team2"`
}

func (g *GameModel) ToGameEntity() *GameEntity {
	mdate := time.Unix(g.Game.Mdate, 0).Format("02 January 2006")

	return &GameEntity{
		Id:      g.Game.Id,
		Mdate:   mdate,
		Stadium: g.Game.Stadium,
		Team1:   g.Game.Team1,
		Team2:   g.Game.Team2,
	}
}
