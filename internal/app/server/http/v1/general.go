package v1

import (
	"github.com/abdil1234/test-golang/internal/app/controllers"
	"github.com/gin-gonic/gin"
)

type GeneralRoute struct {
	routeGroup *gin.RouterGroup
}

func (g *GeneralRoute) Routes(opt controllers.ControllerOption) {
	group := g.routeGroup
	gameController := controllers.NewGameController(opt)

	group.GET("games", gameController.Gets)
	group.POST("game", gameController.Insert)
}

func NewGeneralRoute(rg *gin.RouterGroup) *GeneralRoute {
	return &GeneralRoute{
		routeGroup: rg,
	}
}
