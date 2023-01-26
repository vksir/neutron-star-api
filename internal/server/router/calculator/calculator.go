package calculator

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"neutron-star-api/internal/server/model/calculator/calculatorreq"
	"neutron-star-api/internal/server/model/common/commonresp"
)

func LoadRouters(g *gin.RouterGroup) {
	g.GET("/calculator/ipv4", ipv4Calc)
	g.GET("/calculator/ipv6", ipv6Calc)
}

// ipv4Calc godoc
// @Summary      解析 ipv4
// @Tags         calculator
// @Accept       json
// @Produce      json
// @Param        params  query  calculatorreq.IpCalcParams true "params"
// @Success      200  {object}  calculatorresp.IPv4
// @Failure      500  {object}  commonresp.Err
// @Router       /calculator/ipv4 [get]
func ipv4Calc(c *gin.Context) {
	var params calculatorreq.IpCalcParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, commonresp.Err{Detail: err.Error()})
		return
	}
	if i, err := ParseIPv4(params.IpAddr); err != nil {
		c.JSON(http.StatusBadRequest, commonresp.Err{Detail: err.Error()})
	} else {
		c.JSON(http.StatusOK, i)
	}
}

// ipv6Calc godoc
// @Summary      解析 ipv6
// @Tags         calculator
// @Accept       json
// @Produce      json
// @Param        params  query  calculatorreq.IpCalcParams true "params"
// @Success      200  {object}  calculatorresp.IPv6
// @Failure      500  {object}  commonresp.Err
// @Router       /calculator/ipv6 [get]
func ipv6Calc(c *gin.Context) {
	var params calculatorreq.IpCalcParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, commonresp.Err{Detail: err.Error()})
		return
	}
	if i, err := ParseIPv6(params.IpAddr); err != nil {
		c.JSON(http.StatusBadRequest, commonresp.Err{Detail: err.Error()})
	} else {
		c.JSON(http.StatusOK, i)
	}
}
