package static

import "github.com/gin-gonic/gin"

const (
	Root         = "/opt/neutron_star_api"
	RelativePath = "/static"

	PixivRoot         = "/opt/neutron_star_api/img/pixiv"
	PixivRelativePath = "/static/img/pixiv"
)

func LoadRouters(g *gin.RouterGroup) {
	g.Static("/static", Root)
}
