package pixiv

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
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
// @Success      200  {object}  pixiv.ImageInDB
// @Failure      400  {object}  model.Err
// @Failure      500  {object}  model.Err
// @Router       /pixiv [get]
func get(c *gin.Context) {
	var params GetParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, model.Err{Detail: err.Error()})
		return
	}
	data, err := loliconapi.GetPixivImages(params.Num)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Err{Detail: err.Error()})
		return
	}
	var imgUrls []string
	for _, pixivImg := range data.Data {
		proxyPixivUrl := transformToProxyUrl(pixivImg.Urls.Original)
		imageInDB, isExists := findImageInDB(pixivImg.Pid)
		if isExists {
			localUrl := url.URL{
				Scheme: "http",
				Host:   c.Request.Host,
				Path:   imageInDB.RelativePath,
			}
			imgUrls = append(imgUrls, localUrl.String())
		} else {
			imgUrls = append(imgUrls, proxyPixivUrl)
			go downloadFromPixiv(proxyPixivUrl, pixivImg)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"img_urls": imgUrls,
	})
}
