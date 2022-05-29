package model

type PixivErr struct {
	Detail string `json:"detail"`
}

type ImageInDB struct {
	Id           int
	pid          int
	uid          int
	Tags         string
	IsR18        bool
	OriginUrl    string
	RelativePath string
	CreateTime   string
}

type GetParams struct {
	Num int `form:"num,default=1"`
}
