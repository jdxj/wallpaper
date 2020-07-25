package darenyou

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/astaxie/beego/logs"

	"github.com/jdxj/wallpaper/client"
	"github.com/panjf2000/ants/v2"

	"github.com/PuerkitoBio/goquery"

	"github.com/jdxj/wallpaper/downloader"
	"github.com/jdxj/wallpaper/utils"
)

const (
	//mainPage = "https://darenyouphoto.com/_api/v0/site/youdaren/projects"
	mainPage = "https://darenyouphoto.com/_api/v0/site/youdaren/projects?type=page&offset=0&limit=40"
)

func New(flags *Flags) *DaRenYou {
	pool, _ := ants.NewPool(10)
	c := &DaRenYou{
		gp:    pool,
		c:     client.New(),
		flags: flags,
	}
	return c
}

type DaRenYou struct {
	gp *ants.Pool
	c  *http.Client

	flags *Flags
}

func (dry *DaRenYou) Run() {
	project, err := dry.parseJson()
	if err != nil {
		logs.Error("%s", err)
		return
	}

	urls, err := dry.parseURL(project)
	if err != nil {
		logs.Error("%s", err)
		return
	}

	for _, v := range urls {
		fileName := utils.TruncateFileName(v)
		reqTask := &downloader.RequestTask{
			Path:     dry.cmdParser.path,
			FileName: fileName,
			URL:      v,
		}

		if err := dry.downloader.PushTask(reqTask); err != nil {
			fmt.Printf("PushURL-PushTask err: %s\n", err)
			continue
		}
	}
	dry.downloader.WaitSave()
}

func (dry *DaRenYou) submitTask(downloadLink string) {
	t := &task{}
	if err := dry.gp.Submit(t.Func); err != nil {
		logs.Error("%s", err)
	}
}

func (dry *DaRenYou) parseJson() (*Project, error) {
	c := dry.c
	resp, err := c.Get(mainPage)
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
	switch dry.flags.Project {
	case Chaos:
		project = projects[0]
	case Hysteresis:
		project = projects[1]
	case Commissioned:
		project = projects[2]
	default:
		return nil, fmt.Errorf("don't have this project: %s",
			dry.flags.Project)
	}
	return project, nil
}

func (dry *DaRenYou) parseURL(project *Project) ([]string, error) {
	reader := bytes.NewReader(project.Content)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	var selector string
	switch dry.flags.Size {
	case Src:
		selector = Src
	case SrcO:
		selector = SrcO
	case DataHiRes:
		selector = DataHiRes
	default:
		return nil, fmt.Errorf("no this size: %s", dry.flags.Size)
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
