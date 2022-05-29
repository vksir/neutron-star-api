package swagger

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	_ "neutron-star-api/docs"
)

func LoadHandler(e *gin.Engine) {
	e.GET("/docs", redirect)
	e.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/docs/index.html")
}
