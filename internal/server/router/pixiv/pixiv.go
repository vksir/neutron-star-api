package pixiv

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"neutron-star-api/internal/server/model"
	"neutron-star-api/third_party/loliconapi"
)

func LoadHandler(e *gin.Engine) {
	e.GET("/pixiv", get)
}

// get godoc
// @Summary      获取 pixiv 图片 url
// @Tags         pixiv
// @Accept       json
// @Produce      json
// @Param        params  query  pixiv.GetParams true "params"
// @Success      200  {object}  pixiv.Images
// @Failure      400  {object}  model.Err
// @Failure      500  {object}  model.Err
// @Router       /pixiv [get]
func get(c *gin.Context) {
	var params GetParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, model.Err{Detail: err.Error()})
		return
	}
	pixivImages, err := loliconapi.GetPixivImages(params.Num)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Err{Detail: err.Error()})
		return
	}
	var imgUrls []string
	for _, pixivImg := range pixivImages {
		proxyPixivUrl := transformToProxyUrl(pixivImg.Urls.Original)
		imgUrls = append(imgUrls, proxyPixivUrl)
	}
	c.JSON(http.StatusOK, gin.H{
		"img_urls": imgUrls,
	})
}
