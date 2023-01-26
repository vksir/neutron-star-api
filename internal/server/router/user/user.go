package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"neutron-star-api/internal/database/userdb"
	"neutron-star-api/internal/server/model/common/commonresp"
	"neutron-star-api/internal/server/model/user/userreq"
	"neutron-star-api/internal/server/model/user/userresp"
)

func LoadRouters(publicGroup, privateGroup *gin.RouterGroup) {
	publicGroup.POST("/user/login", login)
	publicGroup.POST("/user/signup", signup)

	privateGroup.GET("/user/me", me)
}

// login godoc
// @Summary      登录
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        data  body  userreq.User true "body"
// @Success      200  {object}  userresp.Token
// @Failure      500  {object}  commonresp.Err
// @Router       /user/login [post]
func login(c *gin.Context) {
	var u userreq.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, fmt.Sprintf("Invalid body: %s", err))
		return
	}
	if err := verifyPassword(u.Username, u.Password); err != nil {
		c.JSON(http.StatusUnauthorized, fmt.Sprintf("Verify password failed: %s", err))
		return
	}
	tokenString, err := createToken(u.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Create token failed: %s", err))
		return
	}
	c.JSON(http.StatusOK, userresp.Token{
		Token: tokenString,
	})
}

// signup godoc
// @Summary      注册
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        data  body  userreq.User true "body"
// @Success      200  {object}  commonresp.Ok
// @Failure      500  {object}  commonresp.Err
// @Router       /user/signup [post]
func signup(c *gin.Context) {
	var u userreq.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, commonresp.Err{Detail: fmt.Sprintf("Invalid body: %s", err)})
		return
	}
	userInDB, _ := userdb.Get(u.Username)
	if userInDB != nil {
		c.JSON(http.StatusBadRequest, "Username has been used")
		return
	}
	encryptedPassword := encryptPassword(u.Password)
	if err := userdb.Set(u.Username, encryptedPassword); err != nil {
		c.JSON(http.StatusInternalServerError, commonresp.Err{Detail: fmt.Sprintf("Set user db failed: %s", err)})
		return
	}
	c.JSON(http.StatusOK, commonresp.Ok{})
}

// me godoc
// @Summary      个人信息
// @Security  	 ApiKeyAuth
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200  {object}  userresp.User
// @Failure      500  {object}  commonresp.Err
// @Router       /user/me [get]
func me(c *gin.Context) {
	username := c.GetHeader("x-username")
	userInDB, err := userdb.Get(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Get user info from db failed")
		return
	}
	c.JSON(http.StatusOK, userresp.User{
		Username: userInDB.Username,
	})
}
