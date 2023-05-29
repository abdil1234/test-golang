package controllers

import (
	"time"

	"github.com/abdil1234/test-golang/internal/app/domain/repositories/models"
	payloads "github.com/abdil1234/test-golang/internal/app/payloads"
	"github.com/gin-gonic/gin"
)

type GameController struct {
	ControllerOption
	gin.Context
}

func NewGameController(opt ControllerOption) *GameController {
	return &GameController{
		ControllerOption: opt,
	}
}

func (ctl *GameController) Insert(c *gin.Context) {

	var err error
	defer func() {
		if err != nil {
			c.JSON(500, err)
			return
		}
	}()

	var payload payloads.CreateGameParam
	if err = ctl.SetPostParams(c, &payload); err != nil {
		return
	}

	mDate, _ := time.Parse("02 January 2006", payload.Mdate)

	game := models.Game{
		Id:      payload.Id,
		Stadium: payload.Stadium,
		Mdate:   mDate.Unix(),
		Team1:   payload.Team1,
		Team2:   payload.Team2,
	}

	err = ctl.Services.Game.Insert(game)
	c.JSON(200, "success")

}

func (ctl *GameController) Gets(c *gin.Context) {

	var err error
	defer func() {
		if err != nil {
			c.JSON(500, err)
			return
		}
	}()

	games := ctl.Services.Game.Gets()
	if len(games) > 0 {
		c.JSON(200, games)
	}
	c.JSON(404, "empty game")
}
