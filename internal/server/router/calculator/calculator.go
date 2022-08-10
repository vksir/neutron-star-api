package calculator

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"neutron-star-api/internal/server/model"
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
// @Param        params  query  calculator.IpCalcParams true "params"
// @Success      200  {object}  calculator.IPv4
// @Failure      400  {object}  model.Err
// @Failure      500  {object}  model.Err
// @Router       /calculator/ipv4 [get]
func ipv4Calc(c *gin.Context) {
	var params model.CalculatorIpv4CalcParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, model.Err{Detail: err.Error()})
		return
	}
	if i, err := ParseIPv4(params.IpAddr); err != nil {
		c.JSON(http.StatusBadRequest, model.Err{Detail: err.Error()})
	} else {
		c.JSON(http.StatusOK, i)
	}
}

// ipv6Calc godoc
// @Summary      解析 ipv6
// @Tags         calculator
// @Accept       json
// @Produce      json
// @Param        params  query  calculator.IpCalcParams true "params"
// @Success      200  {object}  calculator.IPv6
// @Failure      400  {object}  model.Err
// @Failure      500  {object}  model.Err
// @Router       /calculator/ipv6 [get]
func ipv6Calc(c *gin.Context) {
	var params IpCalcParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, model.Err{Detail: err.Error()})
		return
	}
	if i, err := ParseIPv6(params.IpAddr); err != nil {
		c.JSON(http.StatusBadRequest, model.Err{Detail: err.Error()})
	} else {
		c.JSON(http.StatusOK, i)
	}
}

func timeCalc(c *gin.Context) {
	var data struct {
	}
	if err := c.ShouldBindQuery(&data); err != nil {
		c.JSON(http.StatusBadRequest, model.Err{Detail: err.Error()})
		return
	}
}
