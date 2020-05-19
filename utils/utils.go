package utils

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"

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

// WriteToFileReadCloser 读取 r 中所有数据到文件,
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
