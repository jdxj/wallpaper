package octodex

import (
	"net/http"
	"path"
	"sync"

	"github.com/jdxj/wallpaper/client"

	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/logs"
	"github.com/panjf2000/ants/v2"
)

const (
	mainPage       = "https://octodex.github.com"
	downloadPrefix = mainPage
)

func New(flags *Flags) *Octodex {
	pool, _ := ants.NewPool(50)

	c := &Octodex{
		flags: flags,
		c:     client.New(),
		gp:    pool,
		mutex: &sync.Mutex{},
	}
	c.cond = sync.NewCond(c.mutex)
	return c
}

type Octodex struct {
	flags *Flags
	c     *http.Client
	gp    *ants.Pool

	mutex *sync.Mutex
	cond  *sync.Cond
	total int
}

func (oc *Octodex) Run() {
	dls, err := oc.getDownloadLink()
	if err != nil {
		logs.Error("%s", err)
		return
	}

	for _, dl := range dls {
		oc.submitTask(dl)
	}

	// 利用条件锁进行阻塞
	oc.mutex.Lock()
	// 仍有未完成的任务
	for oc.total != 0 {
		oc.cond.Wait()
	}
	oc.mutex.Unlock()

	oc.Stop()
}

func (oc *Octodex) getDownloadLink() ([]string, error) {
	c := oc.c
	resp, err := c.Get(mainPage)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
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
	return downloadURLs, nil
}

func (oc *Octodex) submitTask(downloadLink string) {
	t := &task{
		oc:           oc,
		downloadLink: downloadLink,
		fileName:     path.Base(downloadLink),
	}

	if err := oc.gp.Submit(t.Func); err != nil {
		logs.Error("%s", err)
		return
	}

	logs.Info("task submitted: %s", t.downloadLink)
	oc.add()
}

func (oc *Octodex) add() {
	oc.mutex.Lock()
	oc.total++
	oc.mutex.Unlock()
}

func (oc *Octodex) sub() {
	oc.mutex.Lock()
	oc.total--
	oc.cond.Signal()
	oc.mutex.Unlock()
}

func (oc *Octodex) Stop() {
	oc.c.CloseIdleConnections()
	oc.gp.Release()
	logs.Info("all task finish")
}
