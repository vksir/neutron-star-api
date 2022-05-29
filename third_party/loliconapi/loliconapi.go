package loliconapi

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"strconv"
)

const ApiUrl = "https://api.lolicon.app/setu/v2"

type Response struct {
	Error string     `json:"error"`
	Data  []PixivImg `json:"data"`
}

type PixivImg struct {
	Pid        int      `json:"pid"`
	P          int      `json:"p"`
	Uid        int      `json:"uid"`
	Title      string   `json:"title"`
	Author     string   `json:"author"`
	R18        bool     `json:"r18"`
	Width      int      `json:"width"`
	Height     int      `json:"height"`
	Tags       []string `json:"tags"`
	Ext        string   `json:"ext"`
	UploadDate int64    `json:"uploadDate"`
	Urls       struct {
		Original string `json:"original"`
	} `json:"urls"`
}

func GetPixivImages(num int) (*Response, error) {
	c := resty.New()
	resp, err := c.R().
		SetQueryParams(map[string]string{
			"r18": "2",
			"num": strconv.Itoa(num),
		}).
		Get(ApiUrl)
	if err != nil {
		return nil, err
	}
	data := &Response{}
	if err := json.Unmarshal(resp.Body(), data); err != nil {
		return nil, err
	}
	if data.Error != "" {
		return nil, errors.New(data.Error)
	}
	return data, nil
}
