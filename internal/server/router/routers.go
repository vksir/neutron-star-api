package router

import (
	"github.com/gin-gonic/gin"
	"neutron-star-api/internal/server/router/calculator"
	"neutron-star-api/internal/server/router/pixiv"
	"neutron-star-api/internal/server/router/static"
	"neutron-star-api/internal/server/router/swagger"
)

func LoadRouters(e *gin.Engine) {
	calculator.LoadHandler(e)
	pixiv.LoadHandler(e)
	static.LoadHandler(e)
	swagger.LoadHandler(e)
}
