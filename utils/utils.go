package utils

import (
	"bufio"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
)

var (
	bufferPool = sync.Pool{
		New: func() interface{} {
			return bufio.NewWriter(nil)
		},
	}
)

// WriteFromReadCloser 读取 r 中所有数据到文件,
// 其必须调用 r.Close() 方法.
func WriteFromReadCloser(path, fileName string, r io.ReadCloser) error {
	defer r.Close()

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}

	path = filepath.Join(path, fileName)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	bufW := bufferPool.Get().(*bufio.Writer)
	defer bufferPool.Put(bufW)
	bufW.Reset(file)
	defer bufW.Flush()

	_, err = bufW.ReadFrom(r)
	return err
}

// ReceiveInterrupt 用于接收中断信号, 其必须在新 goroutine 中.
func ReceiveInterrupt() {
	signals := make(chan os.Signal, 2)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
	select {
	case <-signals:
	}
}
