package pixiv

import (
	"strings"
)

const (
	originPixivDomain = "i.pixiv.cat"
	proxyPixivDomain  = "i.pixiv.re"
)

func transformToProxyUrl(pixivUrl string) string {
	return strings.ReplaceAll(pixivUrl, originPixivDomain, proxyPixivDomain)
}
