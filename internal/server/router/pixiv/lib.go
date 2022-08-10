package pixiv

import (
	"fmt"
	"log"
	"neutron-star-api/internal/database"
	"neutron-star-api/internal/server/router/static"
	"neutron-star-api/pkg/util"
	"neutron-star-api/third_party/loliconapi"
	"os"
	"path"
	"strings"
	"time"
)

const (
	originPixivDomain = "i.pixiv.cat"
	proxyPixivDomain  = "i.pixiv.re"
)

func transformToProxyUrl(pixivUrl string) string {
	return strings.ReplaceAll(pixivUrl, originPixivDomain, proxyPixivDomain)
}

func downloadFromPixiv(downloadUrl string, p loliconapi.PixivImg) {
	filename := fmt.Sprintf("%d.%s", p.Pid, p.Ext)
	savePath := path.Join(static.PixivRoot, filename)
	if err := util.DownloadFile(downloadUrl, savePath); err != nil {
		log.Println("download failed: ", err)
	} else {
		if err = saveImageInDB(p); err != nil {
			log.Println("saveImageInDB failed, begin remove file")
			err := os.Remove(savePath)
			if err != nil {
				log.Printf("remove file failed: file=%s, err=%v", savePath, err)
			}
		}
	}
}

func findImageInDB(pid int) (*ImageInDB, bool) {
	var img ImageInDB
	err := database.DB.QueryRow("select * from t_pixiv where picture_id == ?", pid).Scan(&img)
	if err != nil {
		return nil, false
	}
	return &img, true
}

func saveImageInDB(p loliconapi.PixivImg) error {
	filename := fmt.Sprintf("%d.%s", p.Pid, p.Ext)
	filePath := path.Join(static.PixivRelativePath, filename)
	_, err := database.DB.Exec("insert into t_pixiv (picture_id, author_id, tags, is_r18, origin_url, relative_path, create_time) values ($1, $2, $3, $4, $5, $6, $7)", p.Pid, p.Uid, strings.Join(p.Tags, ","), p.R18, p.Urls.Original, filePath, time.Now().String())
	return err
}
