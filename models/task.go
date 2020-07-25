package models

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/jdxj/wallpaper/utils"

	"github.com/astaxie/beego/logs"
)

var (
	ErrNotImage = errors.New("not image")
)

type task struct {
	cl *Crawler

	fileName     string
	downloadLink string
}

func (t *task) runTask() {
	// 自动重试
	retry := t.cl.flags.Retry
	for i := 0; i < retry; i++ {
		err := t.download()
		if err == nil {
			logs.Info("task finished, download link: %s", t.downloadLink)
			t.cl.subOne()
			return
		}
		logs.Error("task retry, download link: [%d]-%s, err: %s",
			i, t.downloadLink, err)
	}
	logs.Error("task failed, download link: %s", t.downloadLink)
}

func (t *task) download() error {
	c := t.cl.c
	savePath := t.cl.flags.SavePath
	resp, err := c.Get(t.downloadLink)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := checkContentType(resp); err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return utils.WriteToFile(savePath, t.fileName, data)
}

func checkContentType(resp *http.Response) error {
	ct := resp.Header.Get("Content-Type")
	if strings.HasPrefix(ct, "image/") {
		return nil
	}
	return ErrNotImage
}
