package util

import (
	"github.com/go-resty/resty/v2"
)

func DownloadFile(downloadUrl, savePath string) error {
	_, err := resty.New().R().SetOutput(savePath).Get(downloadUrl)
	return err
}
