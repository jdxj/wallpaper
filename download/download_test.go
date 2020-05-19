package download

import (
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
