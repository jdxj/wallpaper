package wallhaven

import (
	"fmt"
	"sync"

	"github.com/PuerkitoBio/goquery"

	"github.com/jdxj/wallpaper/client"
	"github.com/jdxj/wallpaper/download"
	"github.com/jdxj/wallpaper/utils"
)

const (
	APIPrefix = "https://wallhaven.cc/search"

	pageURLLimit = 24
)

func NewCrawler(cp *CmdParser) *Crawler {
	c := &Crawler{
		downloader: download.NewDownloader(),
		cmdParser:  cp,
		pageURLs:   make(chan string, pageURLLimit),
	}
	return c
}

type Crawler struct {
	downloader *download.Downloader
	cmdParser  *CmdParser

	pageURLs chan string
}

func (c *Crawler) PushURL() {
	// 该 goroutine 会快速退出
	go c.parsePageURL()

	wg := sync.WaitGroup{}
	for i := 0; i < download.GoroutineLimit; i++ {
		wg.Add(1)
		go func() {
			c.parseURL()
			wg.Done()
		}()
	}

	wg.Wait() // 确保所有 PushTask goroutine 都在 WaitSave goroutine 前
	c.downloader.WaitSave()
}

func (c *Crawler) parsePageURL() {
	if c.cmdParser.url != "" {
		c.parsePageURLFromSpecified()
		return
	}

	c.parsePageURLFromQuery()
}

// parsePageURLFromSpecified 创建一个下载任务.
func (c *Crawler) parsePageURLFromSpecified() {
	c.pageURLs <- c.cmdParser.url
	close(c.pageURLs)
}

func (c *Crawler) parsePageURLFromQuery() {
	query := c.initialQueryURL()
	resp, err := client.Get(query)
	if err != nil {
		fmt.Printf("parsePageURL-Get err: %s\n", err)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Printf("parsePageURL-NewDocumentFromReader err: %s\n", err)
		return
	}

	selection := doc.Find(".preview")
	selection.Each(func(i int, sel *goquery.Selection) {
		attr, ok := sel.Attr("href")
		if !ok {
			return
		}
		c.pageURLs <- attr
	})
	close(c.pageURLs)
}

func (c *Crawler) initialQueryURL() string {
	cp := c.cmdParser

	query := APIPrefix
	categories := fmt.Sprintf("%d%d%d",
		utils.BoolToInt(cp.general),
		utils.BoolToInt(cp.anime),
		utils.BoolToInt(cp.people),
	)
	query = fmt.Sprintf("%s?categories=%s", query, categories)

	purity := fmt.Sprintf("%d%d%d",
		utils.BoolToInt(cp.sfw),
		utils.BoolToInt(cp.sketchy),
		utils.BoolToInt(cp.nsfw),
	)
	query = fmt.Sprintf("%s&purity=%s", query, purity)

	if cp.topRange != "" {
		query = fmt.Sprintf("%s&topRange=%s", query, cp.topRange)
	}

	query = fmt.Sprintf("%s&sorting=%s", query, cp.sorting)
	query = fmt.Sprintf("%s&order=%s", query, cp.order)
	query = fmt.Sprintf("%s?page=%d", query, cp.page)
	return query
}

func (c *Crawler) parseURL() {
	for url := range c.pageURLs {
		resp, err := client.LimitedGet(url)
		if err != nil {
			fmt.Printf("parseURL-LimitedGet err: %s\n", err)
			continue
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			fmt.Printf("parseURL-NewDocumentFromReader err: %s\n", err)
			_ = resp.Body.Close()
			continue
		}

		sel := doc.Find("#wallpaper")
		attr, ok := sel.Attr("src")
		if !ok {
			fmt.Printf("not found wallpaper: %s\n", url)
			_ = resp.Body.Close()
			continue
		}

		fileName := utils.TruncateFileName(attr)
		reqTask := &download.RequestTask{
			Path:     c.cmdParser.path,
			FileName: fileName,
			URL:      attr,
		}

		if err := c.downloader.PushTask(reqTask); err != nil {
			fmt.Printf("parseURL-PushTask err: %s\n", err)
			_ = resp.Body.Close()
			continue
		}
		_ = resp.Body.Close()
	}
}
