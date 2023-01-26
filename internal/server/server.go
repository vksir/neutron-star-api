package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"neutron-star-api/internal/server/middleware"
	"neutron-star-api/internal/server/router/calculator"
	"neutron-star-api/internal/server/router/pixiv"
	"neutron-star-api/internal/server/router/static"
	"neutron-star-api/internal/server/router/swagger"
	"neutron-star-api/internal/server/router/user"
)

func Run() {
	e := gin.Default()
	loadRouters(e)
	err := e.Run("0.0.0.0:5777")
	if err != nil {
		log.Panic(err)
	}
}

func loadRouters(e *gin.Engine) {
	publicGroup := e.Group("")
	privateGroup := e.Group("", middleware.Jwt())

	calculator.LoadRouters(publicGroup)
	pixiv.LoadRouters(publicGroup)
	static.LoadRouters(publicGroup)
	swagger.LoadRouters(publicGroup)
	user.LoadRouters(publicGroup, privateGroup)
}
