package calculator

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"neutron-star-api/internal/server/router/calculator/ip"
	"neutron-star-api/internal/server/router/calculator/model"
)

func LoadHandler(e *gin.Engine) {
	e.GET("/calculator/ipv4", ipv4Calc)
	e.GET("/calculator/ipv6", ipv6Calc)
	e.GET("/calculator/time", timeCalc)
}

// ipv4Calc godoc
// @Summary      解析 ipv4
// @Tags         calculator
// @Accept       json
// @Produce      json
// @Param        params  query  model.Ipv4CalcParams true "params"
// @Success      200  {object}  model.IPv4
// @Failure      400  {object}  model.CalculatorErr
// @Failure      500  {object}  model.CalculatorErr
// @Router       /calculator/ipv4 [get]
func ipv4Calc(c *gin.Context) {
	var params struct {
		IpAddr string `form:"ip_addr" binding:"required"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, model.CalculatorErr{Detail: err.Error()})
		return
	}
	if i, err := ip.ParseIPv4(params.IpAddr); err != nil {
		c.JSON(http.StatusBadRequest, model.CalculatorErr{Detail: err.Error()})
	} else {
		c.JSON(http.StatusOK, i)
	}
}

// ipv6Calc godoc
// @Summary      解析 ipv6
// @Tags         calculator
// @Accept       json
// @Produce      json
// @Param        params  query  model.Ipv6CalcParams true "params"
// @Success      200  {object}  model.IPv6
// @Failure      400  {object}  model.CalculatorErr
// @Failure      500  {object}  model.CalculatorErr
// @Router       /calculator/ipv6 [get]
func ipv6Calc(c *gin.Context) {
	var params struct {
		IpAddr string `form:"ip_addr" binding:"required"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, model.CalculatorErr{Detail: err.Error()})
		return
	}
	if i, err := ip.ParseIPv6(params.IpAddr); err != nil {
		c.JSON(http.StatusBadRequest, model.CalculatorErr{Detail: err.Error()})
	} else {
		c.JSON(http.StatusOK, i)
	}
}

func timeCalc(c *gin.Context) {
	var data struct {
	}
	if err := c.ShouldBindQuery(&data); err != nil {
		c.JSON(http.StatusBadRequest, model.CalculatorErr{Detail: err.Error()})
		return
	}
}
