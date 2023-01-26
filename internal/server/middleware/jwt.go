package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"neutron-star-api/internal/server/model/common/commonresp"
	"neutron-star-api/internal/server/router/user"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("token")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, commonresp.Err{Detail: "Token not found"})
			c.Abort()
			return
		}
		username, err := user.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, commonresp.Err{Detail: "Invalid token"})
			c.Abort()
			return
		}
		c.Request.Header.Set("x-username", username)
		c.Next()
	}
}
