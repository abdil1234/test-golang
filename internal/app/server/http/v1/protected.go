package v1

import (
	"github.com/abdil1234/test-golang/internal/app/controllers"
	"github.com/gin-gonic/gin"
)

type ProtectedRoute struct {
	routeGroup *gin.RouterGroup
}

func (p *ProtectedRoute) Routes(opt controllers.ControllerOption) {
}

func NewProtectedRoute(rg *gin.RouterGroup) *ProtectedRoute {
	return &ProtectedRoute{
		routeGroup: rg,
	}
}
