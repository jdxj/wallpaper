package models

import (
	"net/http"
	"sync"

	"github.com/jdxj/wallpaper/client"
	"github.com/jdxj/wallpaper/utils"

	"github.com/astaxie/beego/logs"
	"github.com/panjf2000/ants/v2"
)

func NewCrawler(flags *CommonFlags, dli DownloadLinkIterator) *Crawler {
	c := client.New(flags.Timeout)
	dli.SetClient(c)
	gp, _ := ants.NewPool(flags.Concurrent)
	mutex := &sync.Mutex{}

	cl := &Crawler{
		gp:         gp,
		flags:      flags,
		unfinished: 0,
		c:          c,
		dli:        dli,
		mutex:      mutex,
		cond:       sync.NewCond(mutex),
	}
	return cl
}

// Crawler 主要负责执行 task.runTask().
// 其核心部分是并发执行, 该功能交给 ants.Pool,
// 另外 Crawler 负责任务统计, 决定何时退出程序.
type Crawler struct {
	c     *http.Client
	gp    *ants.Pool
	flags *CommonFlags

	// 利用条件锁对程序进行阻塞, 直到所有任务完成才退出.
	mutex      *sync.Mutex
	cond       *sync.Cond
	unfinished int // 未完成的任务数

	dli DownloadLinkIterator
}

// Run 的执行步骤是:
//   1. 利用 DownloadLinkIterator 不断获取下载链接, 注意这里不需要并发执行.
//   2. 将下载任务提交, 任务成功提交后将会由 ants.Pool 并发执行.
func (cl *Crawler) Run() {
	dli := cl.dli
	for dli.HasNext() {
		// 由于这里可能会长时间获取下载链接,
		// 所以该位置用于判断是否停止.
		select {
		case <-utils.Stop:
			logs.Info("stop get download links")
			goto skip
		default:
		}

		dls := dli.Next()
		for _, dl := range dls {
			cl.submitTask(dl)
		}
	}

skip:
	cl.mutex.Lock()
	for cl.unfinished != 0 { // 有未完成的任务
		cl.cond.Wait()
	}
	cl.mutex.Unlock()

	cl.stop()
}

func (cl *Crawler) submitTask(dl DownloadLink) {
	// 可能有部分下载链接没有拦截,
	// 所以这里是一个停止点.
	select {
	case <-utils.Stop:
		logs.Info("stop submit, download link: %s",
			dl.URL())
		return
	default:
	}

	t := &task{
		cl:           cl,
		fileName:     dl.FileName(),
		downloadLink: dl.URL(),
	}

	_ = cl.gp.Submit(t.runTask)
	logs.Info("task submitted, download link: %s", dl.URL())
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
