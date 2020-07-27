package wallhaven

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jdxj/wallpaper/models"

	"github.com/astaxie/beego/logs"
)

const (
	APICollection = "https://wallhaven.cc/api/v1/collections/%s/%s?page=%d"
)

func NewWallhavenDLI(flags *Flags) *WallhavenDLI {
	c := &WallhavenDLI{
		flags:  flags,
		ticker: time.NewTicker(1500 * time.Millisecond),
	}
	return c
}

type WallhavenDLI struct {
	c     *http.Client
	flags *Flags

	currPage int
	lastPage int

	// wallhaven 的 api 有访问限制,
	// 但是下载图片的链接没有限制.
	ticker *time.Ticker
}

func (wd *WallhavenDLI) SetClient(c *http.Client) {
	wd.c = c
}

func (wd *WallhavenDLI) HasNext() (has bool) {
	defer func() {
		if !has {
			wd.ticker.Stop()
		}
	}()

	limit := wd.flags.Limit
	if limit > 0 { // 只下载前几页
		return wd.currPage < limit
	}
	return wd.currPage < wd.lastPage
}

func (wd *WallhavenDLI) Next() []models.DownloadLink {
	// 避免过快的访问
	select {
	case <-wd.ticker.C:
	}

	wd.currPage++
	dls, err := wd.parseDownloadLinks(wd.currPage)
	if err != nil {
		logs.Error("%s", err)
		wd.currPage--
		return nil
	}
	return dls
}

func (wd *WallhavenDLI) getQueryURL(page int) string {
	flags := wd.flags
	query := fmt.Sprintf(APICollection, flags.UserName, flags.CollectionID, page)
	if flags.APIKey != "" {
		query = fmt.Sprintf("%s&apikey=%s", query, flags.APIKey)
	}
	return query
}

func (wd *WallhavenDLI) parseDownloadLinks(page int) ([]models.DownloadLink, error) {
	qu := wd.getQueryURL(page)
	resp, err := wd.c.Get(qu)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respJson := &Response{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(respJson); err != nil {
		return nil, err
	}

	images := respJson.Images
	downloadLinks := make([]models.DownloadLink, 0, len(images))
	for _, image := range images {
		dl := &whDL{
			downloadLink: image.Path,
			id:           image.ID,
			purity:       image.Purity,
			category:     image.Category,
		}
		downloadLinks = append(downloadLinks, dl)
	}

	// 更新状态
	meta := respJson.Meta
	wd.lastPage = meta.LastPage
	return downloadLinks, nil
}

type whDL struct {
	downloadLink string

	id       string
	purity   string
	category string
}

func (wd *whDL) URL() string {
	return wd.downloadLink
}

func (wd *whDL) FileName() string {
	return fmt.Sprintf("%s_%s_%s",
		wd.id, wd.purity, wd.category)
}
