package album

import (
	"fmt"
	"net/http"

	"github.com/jdxj/wallpaper/app/androidesk"

	"github.com/astaxie/beego/logs"
)

const (
	APIAlbumList = "http://service.picasso.adesk.com/v1/wallpaper/album?adult=%s&limit=10&skip=%d&order=%s"

	APIAlbum   = "http://service.picasso.adesk.com/v1/wallpaper/album/%s/wallpaper?limit=30&skip=%d&adult=%t&order=%s"
	AlbumLimit = 30

	Host      = "service.picasso.adesk.com"
	UserAgent = "picasso,274,tencent"
)

func NewAlbumDLI(flags *Flags) *AlbumDLI {
	ad := &AlbumDLI{
		flags:   flags,
		curPage: -1,
		amount:  AlbumLimit, // 假设第一次访问的链接有图片
	}
	return ad
}

type AlbumDLI struct {
	c     *http.Client
	flags *Flags

	curPage int
	amount  int
}

func (ad *AlbumDLI) SetClient(c *http.Client) {
	ad.c = c
}

func (ad *AlbumDLI) HasNext() bool {
	limit := ad.flags.Limit
	if limit > 0 {
		// AlbumDLI.curPage 是从0开始,
		// limit 是从1开始.
		return ad.curPage < limit
	}
	return ad.amount >= AlbumLimit
}

func (ad *AlbumDLI) Next() []string {
	ad.curPage++
	qu := ad.getQueryURL(ad.curPage)
	dls, err := ad.parseDownloadLinks(qu)
	if err != nil {
		ad.curPage--
		logs.Error("%s", err)
		return nil
	}
	return dls
}

func (ad *AlbumDLI) getQueryURL(page int) string {
	flags := ad.flags
	return fmt.Sprintf(APIAlbum, flags.ID, ad.pageToSkip(page), flags.Adult, flags.Order)
}

func (ad *AlbumDLI) pageToSkip(page int) int {
	return page * AlbumLimit
}

func (ad *AlbumDLI) parseDownloadLinks(qu string) ([]string, error) {
	resp, err := ad.c.Do(newReq(qu))
	if err != nil {
		return nil, err
	}

	respJson := &androidesk.Response{}
	if err := respJson.Unmarshal(resp.Body); err != nil {
		return nil, err
	}

	w := &Wallpapers{}
	if err := w.Unmarshal(respJson.Resource); err != nil {
		return nil, err
	}

	ad.amount = len(w.Data)
	downloadLinks := make([]string, 0, ad.amount)
	for _, ww := range w.Data {
		downloadLinks = append(downloadLinks, ww.Img)
	}
	if len(downloadLinks) == 0 {
		logs.Warn("not get download links, query url: %s", qu)
	}
	return downloadLinks, nil
}

func newReq(qu string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, qu, nil)
	req.Header.Set("Host", Host)
	req.Header.Set("User-Agent", UserAgent)
	return req
}
