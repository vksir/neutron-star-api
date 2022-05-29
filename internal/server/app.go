package server

import (
	"github.com/gin-gonic/gin"
	"neutron-star-api/internal/server/router"
)

// @title           Neutron Star API
// @version         1.0

func Run() {
	e := gin.Default()
	router.LoadRouters(e)
	e.Run("0.0.0.0:5777")
}
