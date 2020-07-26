package wallhaven

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/astaxie/beego/logs"
)

const (
	APICollection = "https://wallhaven.cc/api/v1/collections/%s/%s?page=%d"
)

func NewWallhavenDLI(flags *Flags) *WallhavenDLI {
	c := &WallhavenDLI{
		flags: flags,
	}
	return c
}

type WallhavenDLI struct {
	c     *http.Client
	flags *Flags

	currPage int
	lastPage int
}

func (wd *WallhavenDLI) SetClient(c *http.Client) {
	wd.c = c
}

func (wd *WallhavenDLI) HasNext() bool {
	limit := wd.flags.Limit
	if limit > 0 { // 只下载前几页
		return wd.currPage <= limit
	}
	return wd.currPage <= wd.lastPage
}

func (wd *WallhavenDLI) Next() []string {
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

func (wd *WallhavenDLI) parseDownloadLinks(page int) ([]string, error) {
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
	downloadLinks := make([]string, 0, len(images))
	for _, image := range images {
		downloadLinks = append(downloadLinks, image.Path)
	}

	// 更新状态
	meta := respJson.Meta
	wd.lastPage = meta.LastPage
	return downloadLinks, nil
}
