package octodex

import (
	"encoding/json"
	"fmt"

	"github.com/jdxj/wallpaper/cache"
	"github.com/jdxj/wallpaper/client"
	"github.com/jdxj/wallpaper/download"
	"github.com/jdxj/wallpaper/utils"

	"github.com/PuerkitoBio/goquery"
)

const (
	mainPage       = "https://octodex.github.com"
	downloadPrefix = mainPage
)

func NewCrawler(flags *Flags) *Crawler {
	c := &Crawler{
		downloader: download.NewDownloader(),
		flags:      flags,
	}
	return c
}

type Crawler struct {
	downloader *download.Downloader
	flags      *Flags
}

// PushURL 不断地获取下载链接
func (oc *Crawler) PushURL() {
	result, err := cache.IsVisited(key)
	if err != nil {
		fmt.Printf("PushURL-IsVisited err: %s\n", err)

		oc.pushURLFromWeb()
		return
	}

	du := &downloadURLs{}
	if err := json.Unmarshal(result, du); err != nil {
		fmt.Printf("PushURL-Unmarshal err: %s\n", err)

		oc.pushURLFromWeb()
		return
	}

	oc.pushURLFromCache(du.Urls)
	oc.downloader.WaitSave()
}

func (oc *Crawler) pushURLFromCache(urls []string) {
	fmt.Printf("get urls from cache!\n")
	for _, url := range urls {
		oc.pushURL(url)
	}
}

func (oc *Crawler) pushURLFromWeb() {
	fmt.Printf("get urls from web!\n")
	resp, err := client.Get(mainPage)
	if err != nil {
		fmt.Printf("PushURL-Get err: %s\n", err)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Printf("PushURL-NewDocumentFromReader err: %s\n", err)
		return
	}
	// 缓存用
	var urls []string
	sel := doc.Find(".width-fit")
	sel.Each(func(i int, selI *goquery.Selection) {
		// src 的格式为: "/images/Octoqueer.png"
		src, ok := selI.Attr("data-src")
		if !ok {
			return
		}
		url := downloadPrefix + src
		urls = append(urls, url)
		oc.pushURL(url)
	})

	cacheUrls(urls)
}

func (oc *Crawler) pushURL(url string) {
	fileName := utils.TruncateFileName(url)
	reqTask := &download.RequestTask{
		Path:     oc.flags.Path,
		FileName: fileName,
		URL:      url,
	}

	if err := oc.downloader.PushTask(reqTask); err != nil {
		fmt.Printf("pushURL-PushTask err: %s\n", err)
	}
}

func cacheUrls(urls []string) {
	du := &downloadURLs{
		Urls: urls,
	}
	data, _ := json.Marshal(du)
	if err := cache.SaveValue(key, data); err != nil {
		fmt.Printf("pushURLFromWeb-SaveValue: err: %s\n", err)
	}
}
