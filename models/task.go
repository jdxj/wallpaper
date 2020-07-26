package models

import (
	"errors"
	"net/http"
	"strings"

	"github.com/jdxj/wallpaper/utils"

	"github.com/astaxie/beego/logs"
)

var (
	ErrNotImage            = errors.New("not image")
	ErrContentTypeNotFound = errors.New("Content-Type not found")
)

type task struct {
	cl *Crawler

	fileName     string
	downloadLink string
}

func (t *task) runTask() {
	// 对于提交成功的任务, 不管下载成功还是失败,
	// 下载结束后都要调用 subOne().
	defer t.cl.subOne()

	// 自动重试
	retry := t.cl.flags.Retry
	for i := 0; i < retry; i++ {
		// 这里是个耗时操作, 提前拦截
		select {
		case <-utils.Stop:
			logs.Info("stop runTask, download link: %s",
				t.downloadLink)
			return
		default:
		}

		// 进入下载的任务不中断
		err := t.download()
		if err == nil {
			logs.Info("task finished, download link: %s", t.downloadLink)
			return
		}

		logs.Error("task retry, download link: [%d]-%s, err: %s",
			i, t.downloadLink, err)
		if err == ErrNotImage {
			// 如果不是 image, 则直接退出而不重试.
			break
		}
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

	return utils.WriteFromReadCloser(savePath, t.fileName, resp.Body)
}

func checkContentType(resp *http.Response) error {
	ct := resp.Header.Get("Content-Type")
	if ct == "" {
		return ErrContentTypeNotFound
	}
	if strings.HasPrefix(ct, "image/") {
		return nil
	} else if strings.HasSuffix(ct, "octet-stream") {
		return nil
	}
	return ErrNotImage
}
