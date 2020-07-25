package octodex

import (
	"net/http"

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

func (oDLI *OctodexDLI) Next() []string {
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

	downloadURLs := make([]string, 0, 200)
	sel := doc.Find(".width-fit")
	sel.Each(func(i int, selI *goquery.Selection) {
		// src 的格式为: "/images/Octoqueer.png"
		src, ok := selI.Attr("data-src")
		if !ok {
			return
		}
		url := downloadPrefix + src
		downloadURLs = append(downloadURLs, url)
	})
	oDLI.hasNext = false
	return downloadURLs
}
