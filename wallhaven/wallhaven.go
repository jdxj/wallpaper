package wallhaven

import (
	"fmt"
	"sync"

	"github.com/jdxj/wallpaper/cache"

	"github.com/PuerkitoBio/goquery"

	"github.com/jdxj/wallpaper/downloader"
	"github.com/jdxj/wallpaper/utils"
)

const (
	APIPrefix = "https://wallhaven.cc/search"

	pageURLLimit = 24
)

func NewCrawler(flags *Flags) *Crawler {
	c := &Crawler{
		downloader: downloader.NewDownloader(),
		flags:      flags,
		pageURLs:   make(chan string, pageURLLimit),
	}
	return c
}

type Crawler struct {
	downloader *downloader.Downloader
	flags      *Flags

	pageURLs chan string
}

func (c *Crawler) PushURL() {
	// 该 goroutine 会快速退出
	go c.parsePageURL()

	wg := sync.WaitGroup{}
	for i := 0; i < downloader.GoroutineLimit; i++ {
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
	if c.flags.Url != "" {
		c.parsePageURLFromSpecified()
		return
	}

	c.parsePageURLFromQuery()
}

// parsePageURLFromSpecified 创建一个下载任务.
func (c *Crawler) parsePageURLFromSpecified() {
	c.pageURLs <- c.flags.Url
	close(c.pageURLs)
}

func (c *Crawler) parsePageURLFromQuery() {
	query := c.initialQueryURL()
	resp, err := downloader.Get(query)
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
	flags := c.flags

	query := APIPrefix
	categories := fmt.Sprintf("%d%d%d",
		utils.BoolToInt(flags.General),
		utils.BoolToInt(flags.Anime),
		utils.BoolToInt(flags.People),
	)
	query = fmt.Sprintf("%s?categories=%s", query, categories)

	purity := fmt.Sprintf("%d%d%d",
		utils.BoolToInt(flags.Sfw),
		utils.BoolToInt(flags.Sketchy),
		utils.BoolToInt(flags.Nsfw),
	)
	query = fmt.Sprintf("%s&purity=%s", query, purity)

	if flags.TopRange != "" {
		query = fmt.Sprintf("%s&topRange=%s", query, flags.TopRange)
	}

	query = fmt.Sprintf("%s&sorting=%s", query, flags.Sorting)
	query = fmt.Sprintf("%s&order=%s", query, flags.Order)
	query = fmt.Sprintf("%s&page=%d", query, flags.Page)
	return query
}

func (c *Crawler) parseURL() {
	for url := range c.pageURLs {
		// 查询缓存
		// todo: 实现
		value, err := cache.IsVisited(nil)
		if err != nil || value == nil { // 未命中
			fmt.Printf("parseURL-check cache faild, value: %s, err: %s\n",
				value, err)
		} else { // 命中
			fmt.Printf("parseURL-IsVisited, hit cache-> key: %s, value: %s\n",
				url, value)
			c.pushTask("") // todo
			continue
		}
		// 未命中后进行 http 访问
		imgURL, err := c.getImgURL(url)
		if err != nil {
			fmt.Printf("parseURL-getImgURL err: %s", err)
			continue
		}

		c.pushTask(imgURL)
		// 进行缓存
		// todo
		if err := cache.SaveValue(nil, nil); err != nil {
			fmt.Printf("parseURL-SaveValue err: %s\n", err)
		}
	}
}

func (c *Crawler) getImgURL(preURL string) (string, error) {
	resp, err := downloader.LimitedGet(preURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	sel := doc.Find("#wallpaper")
	attr, ok := sel.Attr("src")
	if !ok {
		return "", fmt.Errorf("not found wallpaper: %s\n", preURL)
	}
	return attr, nil
}

func (c *Crawler) pushTask(url string) {
	fileName := utils.TruncateFileName(url)
	reqTask := &downloader.RequestTask{
		Path:     c.flags.Path,
		FileName: fileName,
		URL:      url,
	}

	if err := c.downloader.PushTask(reqTask); err != nil {
		fmt.Printf("pushTask-PushTask err: %s\n", err)
	}
}
