package darenyou

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/jdxj/wallpaper/utils"

	"github.com/jdxj/wallpaper/client"
	"github.com/jdxj/wallpaper/download"
)

const (
	//mainPage = "https://darenyouphoto.com/_api/v0/site/youdaren/projects"
	mainPage = "https://darenyouphoto.com/_api/v0/site/youdaren/projects?type=page&offset=0&limit=40"

	savePath = "data"

	// project
	Chaos        = "chaos"
	Commissioned = "commissioned"

	// size
	Src       = "src"
	SrcO      = "src_o"
	DataHiRes = "data-hi-res" // goquery 无法解析
)

func NewCrawler(project, size string) *Crawler {
	c := &Crawler{
		downloader: download.NewDownloader(),
		project:    project,
		size:       size,
	}
	return c
}

type Crawler struct {
	downloader *download.Downloader

	project string
	size    string
}

func (c *Crawler) PushURL() {
	project, err := c.parseJson()
	if err != nil {
		fmt.Printf("PushURL-parseJson err: %s\n", err)
		return
	}

	urls, err := c.parseURL(project)
	if err != nil {
		fmt.Printf("PushURL-parseURL err: %s\n", err)
		return
	}
	fmt.Printf("urls len: %d\n", len(urls))

	for _, v := range urls {
		fileName := utils.TruncateFileName(v)
		reqTask := &download.RequestTask{
			Path:     savePath,
			FileName: fileName,
			URL:      v,
		}

		if err := c.downloader.PushTask(reqTask); err != nil {
			fmt.Printf("PushURL-PushTask err: %s\n", err)
			continue
		}
	}
	c.downloader.WaitSave()
}

func (c *Crawler) parseJson() (*Project, error) {
	resp, err := client.Get(mainPage)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	projects := make([]*Project, 0)
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&projects); err != nil {
		return nil, err
	}

	if len(projects) < 4 {
		return nil, fmt.Errorf("project num not enough: %d",
			len(projects))
	}

	var project *Project
	switch c.project {
	case Chaos:
		project = projects[0]

	case Commissioned:
		project = projects[1]

	default:
		return nil, fmt.Errorf("don't have this project: %s",
			c.project)
	}
	return project, nil
}

func (c *Crawler) parseURL(project *Project) ([]string, error) {
	reader := bytes.NewReader(project.Content)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	var selector string
	switch c.size {
	case Src:
		selector = Src

	case SrcO:
		selector = SrcO

	case DataHiRes:
		selector = DataHiRes

	default:
		return nil, fmt.Errorf("no this size: %s", c.size)
	}

	var result []string
	sel := doc.Find("img")
	sel.Each(func(i int, selection *goquery.Selection) {
		attr, ok := selection.Attr(selector)
		if !ok {
			return
		}

		// 清除无用字符
		attr = strings.ReplaceAll(attr, `\`, "")
		attr = strings.ReplaceAll(attr, `"`, "")
		result = append(result, attr)
	})
	return result, nil
}
