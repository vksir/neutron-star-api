package pixiv

type Images struct {
	ImgUrls []string `json:"img_urls"`
}

type GetParams struct {
	Num int `form:"num,default=1"`
}
