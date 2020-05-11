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

func WriteToFileReadCloser(path string, r io.ReadCloser) error {
	defer r.Close()

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	bufW := bufio.NewWriter(file)
	defer bufW.Flush()

	_, err = bufW.ReadFrom(r)
	return err
}
