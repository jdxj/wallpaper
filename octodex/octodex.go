package octodex

import (
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/jdxj/wallpaper/client"
	"github.com/jdxj/wallpaper/download"
	"github.com/jdxj/wallpaper/utils"
)

func Run() {
	start := time.Now()
	oc := NewCrawler()
	oc.PushURL()
	end := time.Now()

	fmt.Printf("start:  %s\n", start)
	fmt.Printf("end:    %s\n", end)
	fmt.Printf("expend: %s\n", end.Sub(start))
}

const (
	mainPage       = "https://octodex.github.com"
	downloadPrefix = mainPage
	cacheLimit     = 100
	savePath       = "data"
)

func NewCrawler() *Crawler {
	c := &Crawler{
		downloader: download.NewDownloader(),
	}
	return c
}

type Crawler struct {
	downloader *download.Downloader
}

// PushURL 不断地获取下载链接
func (oc *Crawler) PushURL() {
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

	sel := doc.Find(".width-fit")
	sel.Each(func(i int, selI *goquery.Selection) {
		// src 的格式为: "/images/Octoqueer.png"
		src, ok := selI.Attr("data-src")
		if !ok {
			return
		}

		fileName := utils.TruncateFileName(src)
		reqTask := &download.RequestTask{
			Path:     savePath,
			FileName: fileName,
			URL:      downloadPrefix + "/" + src,
		}
		if err := oc.downloader.PushTask(reqTask); err != nil {
			fmt.Printf("PushURL-PushTask err: %s\n", err)
		}
	})

	oc.downloader.WaitSave()
}
