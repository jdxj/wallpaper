package download

import (
	"fmt"
	"io"
	"sync"

	"github.com/jdxj/wallpaper/client"
	"github.com/jdxj/wallpaper/utils"
)

const (
	GoroutineLimit = 4
	RequestLimit   = 4
	SaveLimit      = RequestLimit * 2
)

// RequestTask 保存了要下载的任务
type RequestTask struct {
	Path     string
	FileName string
	URL      string
}

// saveTask 保存了获取数据的 io
type saveTask struct {
	reqTask *RequestTask
	data    io.ReadCloser
}

func NewDownloader() *Downloader {
	return NewDownloaderSize(RequestLimit, SaveLimit)
}

func NewDownloaderSize(requestLimit, saveLimit int) *Downloader {
	d := &Downloader{
		reqTasks:  make(chan *RequestTask, requestLimit),
		saveTasks: make(chan *saveTask, saveLimit),
		stop:      make(chan int),
		giveUp:    make(chan int),
		reqWG:     &sync.WaitGroup{},
		saveWG:    &sync.WaitGroup{},
	}

	d.processingTasks()
	d.GiveUp()
	return d
}

// Downloader 将负责下载图片文件
type Downloader struct {
	reqTasks chan *RequestTask
	reqWG    *sync.WaitGroup

	saveTasks chan *saveTask
	saveWG    *sync.WaitGroup

	stop   chan int
	giveUp chan int
}

// PushTask 将 RequestTask 存入缓存中.
func (d *Downloader) PushTask(requestTask *RequestTask) error {
	select {
	case <-d.stop:
		return fmt.Errorf("downloader already closed push task channel")

	case <-d.giveUp:
		return fmt.Errorf("downloader is giving up on push task channel")

	default:
	}

	fmt.Printf("pushing: %s\n", requestTask.FileName)
	d.reqTasks <- requestTask
	return nil
}

// processingTasks 负责启动 Downloader 进行下载任务.
func (d *Downloader) processingTasks() {
	reqWG := d.reqWG
	saveWG := d.saveWG

	for i := 0; i < GoroutineLimit; i++ {
		reqWG.Add(1)
		go func() {
			d.getData()
			reqWG.Done()
		}()

		saveWG.Add(1)
		go func() {
			d.saveData()
			saveWG.Done()
		}()
	}
}

// getData 使用提供的 url 来获取 ReadCloser.
func (d *Downloader) getData() {
	for reqTask := range d.reqTasks {
		select {
		case <-d.giveUp:
			fmt.Printf("getData is giving up on: %#v\n", *reqTask)
			continue

		default:
		}

		fmt.Printf("getting: %s\n", reqTask.FileName)
		readCloser, err := client.GetReadCloser(reqTask.URL)
		if err != nil {
			fmt.Printf("getData-GetReadCloser err: %s\n",
				err)
			continue
		}

		saveTask := &saveTask{
			reqTask: reqTask,
			data:    readCloser,
		}
		d.saveTasks <- saveTask
	}
}

// saveData 通过读取 ReadCloser 将数据保存到磁盘.
func (d *Downloader) saveData() {
	for saveTask := range d.saveTasks {
		select {
		case <-d.giveUp:
			fmt.Printf("saveData is giving up on: %s\n",
				saveTask.reqTask.FileName)
			saveTask.data.Close()
			continue

		default:
		}

		path := saveTask.reqTask.Path
		fileName := saveTask.reqTask.FileName

		fmt.Printf("saveing: %s\n", fileName)
		err := utils.WriteFromReadCloser(path, fileName, saveTask.data)
		if err != nil {
			fmt.Printf("saveData-WriteFromReadCloser err: %s\n",
				err)
		}
	}
}

// WaitSave 通知 Downloader 不会再有 RequestTask 被发送至 Downloader.reqTasks
// WaitSave 是阻塞的来等待所有数据保存到磁盘.
// 注意: WaitSave 必须与 PushTask 在同一 goroutine,
//     否则可能会出现 "panic: send on closed channel" 恐慌.
//     或者使用其他手段保证 WaitSave 永远在 PushTask 后运行.
// 注意: WaitSave 必须被调用, 否则只会保存部分数据.
func (d *Downloader) WaitSave() {
	close(d.stop)

	close(d.reqTasks)
	d.reqWG.Wait()

	close(d.saveTasks)
	d.saveWG.Wait()
}

// giveUp 接收中断信号, 放弃缓存立即结束下载.
func (d *Downloader) GiveUp() {
	go func() {
		utils.ReceiveInterrupt()

		fmt.Printf("downloader receive giveup signal\n")
		close(d.giveUp)
	}()
}
