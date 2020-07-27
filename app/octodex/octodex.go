package octodex

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/jdxj/wallpaper/models"

	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/logs"
)

const (
	mainPage       = "https://octodex.github.com"
	downloadPrefix = mainPage
)

func NewOctodexDLI() *OctodexDLI {
	oDLI := &OctodexDLI{
		hasNext: true,
	}
	return oDLI
}

type OctodexDLI struct {
	// client 由 SetClient 设置
	c       *http.Client
	hasNext bool
}

func (oDLI *OctodexDLI) SetClient(client *http.Client) {
	oDLI.c = client
}

func (oDLI *OctodexDLI) HasNext() bool {
	return oDLI.hasNext
}

func (oDLI *OctodexDLI) Next() []models.DownloadLink {
	c := oDLI.c
	resp, err := c.Get(mainPage)
	if err != nil {
		logs.Error("%s", err)
		return nil
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		logs.Error("%s", err)
		return nil
	}

	downloadURLs := make([]models.DownloadLink, 0, 200)
	sel := doc.Find(".width-fit")
	sel.Each(func(i int, selI *goquery.Selection) {
		// src 的格式为: "/images/Octoqueer.png"
		src, ok := selI.Attr("data-src")
		if !ok {
			return
		}
		url := downloadPrefix + src
		dl := &octodexDL{
			downloadLink: url,
		}
		downloadURLs = append(downloadURLs, dl)
	})
	oDLI.hasNext = false
	return downloadURLs
}

type octodexDL struct {
	downloadLink string
}

func (od *octodexDL) URL() string {
	return od.downloadLink
}

func (od *octodexDL) FileName() string {
	suffix := filepath.Base(od.downloadLink)
	idx := strings.LastIndex(suffix, ".")
	if idx < 0 {
		return fmt.Sprintf("octodex_%d", time.Now().UnixNano())
	}
	return suffix[:idx]
}
