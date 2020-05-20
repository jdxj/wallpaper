package utils

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/jdxj/wallpaper/client"
)

func Download(url string) ([]byte, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func WriteToFile(path string, data []byte) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}

// WriteFromReadCloser 读取 r 中所有数据到文件,
// 其必须调用 r.Close() 方法.
func WriteFromReadCloser(path, fileName string, r io.ReadCloser) error {
	defer r.Close()

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}

	file, err := os.Create(path + "/" + fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	bufW := bufio.NewWriter(file)
	defer bufW.Flush()

	_, err = bufW.ReadFrom(r)
	return err
}

// TruncateFileName 用于将 url 中最后一个 "/" 之后的部分截取出来作为文件名.
func TruncateFileName(url string) string {
	idx := strings.LastIndex(url, "/")
	return url[idx+1:]
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// ReceiveInterrupt 用于接收中断信号, 其必须在新 goroutine 中.
func ReceiveInterrupt() {
	signals := make(chan os.Signal, 2)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
	select {
	case <-signals:
	}
}
