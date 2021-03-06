package utils

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"

	"github.com/astaxie/beego/logs"
)

var (
	bufferPool = sync.Pool{
		New: func() interface{} {
			return bufio.NewReader(nil)
		},
	}

	// 停止信号
	Stop = make(chan int)
)

// WriteFromReadCloser 读取 r 中所有数据到文件,
// 其必须调用 r.Close() 方法.
func WriteFromReadCloser(path, fileName string, r io.ReadCloser) error {
	defer r.Close()

	bufR := bufferPool.Get().(*bufio.Reader)
	defer bufferPool.Put(bufR)

	bufR.Reset(r)
	head, err := bufR.Peek(512)
	if err != nil {
		return err
	}
	fileName += GetFileExtension(head)

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}
	path = filepath.Join(path, fileName)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	defer file.Sync()

	_, err = bufR.WriteTo(file)
	return err
}

// ReceiveInterrupt 用于接收中断信号, 其必须在新 goroutine 中.
func ReceiveInterrupt() {
	signals := make(chan os.Signal, 2)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
	select {
	case <-signals:
	}

	logs.Warn("receive stop signal")
	close(Stop)
}

func GetFileExtension(head []byte) string {
	mime := http.DetectContentType(head)

	switch mime {
	case "image/png":
		return ".png"
	case "image/jpeg":
		return ".jpg"
	default:
		return ".unknown"
	}
}
