package http

import (
	v1 "github.com/abdil1234/test-golang/internal/app/server/http/v1"

	"github.com/abdil1234/test-golang/internal/app/controllers"
	"github.com/abdil1234/test-golang/internal/app/server/http/middlewares"
	"github.com/gin-gonic/gin"
)

func Router(opt controllers.ControllerOption) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middlewares.RequestID())

	serviceGroup := r.Group("/golang-test")
	{

		v1Group := serviceGroup.Group("api/v1")
		{
			generalV1Route := v1.NewGeneralRoute(v1Group)
			generalV1Route.Routes(opt)

			protectedV1Route := v1.NewProtectedRoute(v1Group)
			protectedV1Route.Routes(opt)
		}
	}

	return r
}
