package octodex

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/astaxie/beego/logs"

	"github.com/jdxj/wallpaper/utils"
)

var (
	errNotImage = errors.New("not image")
)

type task struct {
	oc    *Octodex
	retry int

	fileName     string
	downloadLink string
}

func (t *task) Func() {
	// 自动重试
	retryLimit := t.oc.flags.Retry
	for i := 0; i < retryLimit; i++ {
		err := t.f()
		// 任务执行成功
		if err == nil {
			logs.Info("task finish: %s", t.downloadLink)
			// 通知主 goroutine
			t.oc.sub()
			return
		}
		logs.Error("task fail: 0-%s, err: %s", i, t.downloadLink, err)
	}
}

func (t *task) f() error {
	c := t.oc.c
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

	path := t.oc.flags.Path
	return utils.WriteToFile(path, t.fileName, data)
}

func checkContentType(resp *http.Response) error {
	ct := resp.Header.Get("Content-Type")
	if strings.HasPrefix(ct, "image/") {
		return nil
	}
	return errNotImage
}
