package darenyou

import (
	"bytes"
	"encoding/json"
	"errors"
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
	//mainPage = "https://darenyouphoto.com/_api/v0/site/youdaren/projects"
	mainPage      = "https://darenyouphoto.com/_api/v0/site/youdaren/projects?type=page&offset=0&limit=40"
	projectAmount = 5
)

var (
	ErrProjectAmountNotEnough = errors.New("project amount not enough")
	ErrProjectNotFound        = errors.New("project not found")
)

func NewDaRenYouDLI(flags *Flags) *DaRenYouDLI {
	c := &DaRenYouDLI{
		flags:   flags,
		hasNext: true,
	}
	return c
}

type DaRenYouDLI struct {
	c     *http.Client
	flags *Flags

	hasNext bool
}

func (dry *DaRenYouDLI) SetClient(c *http.Client) {
	dry.c = c
}

func (dry *DaRenYouDLI) HasNext() bool {
	return dry.hasNext
}

func (dry *DaRenYouDLI) Next() []models.DownloadLink {
	project, err := dry.parseJson()
	if err != nil {
		logs.Error("%s", err)
		return nil
	}

	urls, err := dry.parseURL(project)
	if err != nil {
		logs.Error("%s", err)
		return nil
	}
	dry.hasNext = false
	return urls
}

func (dry *DaRenYouDLI) parseJson() (*Project, error) {
	c := dry.c
	resp, err := c.Get(mainPage)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// projects 相当于不同相册的集合
	projects := make([]*Project, 0)
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&projects); err != nil {
		return nil, err
	}

	if len(projects) > projectAmount {
		logs.Warn("may have created a new project")
	}
	if len(projects) < projectAmount {
		return nil, ErrProjectAmountNotEnough
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
		return nil, ErrProjectNotFound
	}
	return project, nil
}

func (dry *DaRenYouDLI) parseURL(project *Project) ([]models.DownloadLink, error) {
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

	var result []models.DownloadLink
	sel := doc.Find("img")
	sel.Each(func(i int, selection *goquery.Selection) {
		attr, ok := selection.Attr(selector)
		if !ok {
			return
		}
		// 清除无用字符
		attr = strings.ReplaceAll(attr, `\`, "")
		attr = strings.ReplaceAll(attr, `"`, "")

		dl := &dryDL{
			downloadLink: attr,
		}
		result = append(result, dl)
	})
	return result, nil
}

type dryDL struct {
	downloadLink string
}

func (dd *dryDL) URL() string {
	return dd.downloadLink
}

func (dd *dryDL) FileName() string {
	suffix := filepath.Base(dd.downloadLink)
	idx := strings.LastIndex(suffix, "_")
	if idx < 0 {
		return fmt.Sprintf("darenyou_%d", time.Now().UnixNano())
	}
	return suffix[:idx]
}
