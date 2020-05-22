package download

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSendToChan(t *testing.T) {
	c := make(chan int, 100)
	go func() {
		for {
			time.Sleep(time.Second)
			c <- 1
		}
	}()

	time.Sleep(3 * time.Second)
	close(c)
	time.Sleep(10 * time.Second)
}

func TestNewDownloader(t *testing.T) {
	reqTask := &RequestTask{
		Path:     "data",
		FileName: "myPhoto.jpg",
		URL:      "https://static001.geekbang.org/resource/image/7d/50/7dd3f821b4e790b9367af9f23a7e8750.jpg",
	}

	d := NewDownloader()
	_ = d.PushTask(reqTask)

	d.WaitSave()
}

func TestDownloader_PushTask(t *testing.T) {
	reqTask := &RequestTask{
		Path:     "data",
		FileName: "myPhoto.jpg",
		URL:      "https://static001.geekbang.org/resource/image/7d/50/7dd3f821b4e790b9367af9f23a7e8750.jpg",
	}

	d := &Downloader{
		reqTasks:   make(chan *RequestTask, RequestLimit),
		saveTasks:  make(chan *saveTask, SaveLimit),
		stopPush:   make(chan int),
		giveUpSign: make(chan int),
		reqWG:      &sync.WaitGroup{},
		saveWG:     &sync.WaitGroup{},
	}
	// 测试 PushTask 与 WaitSave 在不同 goroutine 中的执行情况
	go func() {
		for i := 0; i < RequestLimit*2; i++ {
			if err := d.PushTask(reqTask); err != nil { // panic: send on closed channel
				fmt.Printf("err: %s\n", err)
				continue
			}
		}
	}()

	time.Sleep(500 * time.Millisecond)
	close(d.reqTasks)
}
