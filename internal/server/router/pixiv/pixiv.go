package pixiv

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"neutron-star-api/internal/server/model/common/commonresp"
	"neutron-star-api/internal/server/model/pixiv/pixivreq"
	"neutron-star-api/internal/server/model/pixiv/pixivresp"
	"neutron-star-api/thirdparty/loliconapi"
)

func LoadRouters(g *gin.RouterGroup) {
	g.GET("/pixiv", get)
}

// get godoc
// @Summary      获取 pixiv 图片 url
// @Tags         pixiv
// @Accept       json
// @Produce      json
// @Param        params  query  pixivreq.GetParams true "params"
// @Success      200  {object}  pixivresp.Images
// @Failure      500  {object}  commonresp.Err
// @Router       /pixiv [get]
func get(c *gin.Context) {
	var params pixivreq.GetParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, commonresp.Err{Detail: err.Error()})
		return
	}
	pixivImages, err := loliconapi.GetPixivImages(params.Num)
	if err != nil {
		c.JSON(http.StatusBadRequest, commonresp.Err{Detail: err.Error()})
		return
	}
	var imgUrls []string
	for _, pixivImg := range pixivImages {
		proxyPixivUrl := transformToProxyUrl(pixivImg.Urls.Original)
		imgUrls = append(imgUrls, proxyPixivUrl)
	}
	c.JSON(http.StatusOK, pixivresp.Images{
		ImgUrls: imgUrls,
	})
}
