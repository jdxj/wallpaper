package models

import (
	"net/http"
	"path"
	"sync"

	"github.com/jdxj/wallpaper/client"

	"github.com/astaxie/beego/logs"
	"github.com/panjf2000/ants/v2"
)

func NewCrawler(cfg *Config, dli DownloadLinkIterator) *Crawler {
	c := client.New()
	dli.SetClient(c)
	gp, _ := ants.NewPool(10)
	mutex := &sync.Mutex{}

	cl := &Crawler{
		gp:         gp,
		cfg:        cfg,
		unfinished: 0,
		c:          c,
		dli:        dli,
		mutex:      mutex,
		cond:       sync.NewCond(mutex),
	}
	return cl
}

type Crawler struct {
	c   *http.Client
	gp  *ants.Pool
	cfg *Config

	mutex      *sync.Mutex
	cond       *sync.Cond
	unfinished int // 未完成的任务数

	dli DownloadLinkIterator
}

func (cl *Crawler) Run() {
	dli := cl.dli
	for dli.HasNext() {
		dls := dli.Next()
		for _, dl := range dls {
			cl.submitTask(dl)
		}
	}

	cl.mutex.Lock()
	for cl.unfinished != 0 { // 有未完成的任务
		cl.cond.Wait()
	}
	cl.mutex.Unlock()

	cl.stop()
}

func (cl *Crawler) submitTask(downloadLink string) {
	t := &Task{
		cl:           cl,
		fileName:     path.Base(downloadLink),
		downloadLink: downloadLink,
	}

	_ = cl.gp.Submit(t.RunTask)
	logs.Info("task submitted, download link: %s", downloadLink)
	cl.addOne()
}

// addOne 表示一个新任务被添加
func (cl *Crawler) addOne() {
	cl.mutex.Lock()
	cl.unfinished++
	cl.mutex.Unlock()
}

// subOne 表示已完成一个任务.
// 如果某个任务已完成, 那么 subOne 负责通知主 goroutine 进行状态检测,
// 如果检测到所有任务都完成, 那么进入结束流程.
func (cl *Crawler) subOne() {
	cl.mutex.Lock()
	cl.unfinished--
	cl.cond.Signal()
	cl.mutex.Unlock()
}

func (cl *Crawler) stop() {
	cl.c.CloseIdleConnections()
	cl.gp.Release()
	logs.Info("all task finished")
}
