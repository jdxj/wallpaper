package downloader

import (
	"github.com/panjf2000/ants/v2"
)

const (
	size = 50
)

func New() (*Downloader, error) {
	pool, err := ants.NewPool(size)
	if err != nil {
		return nil, err
	}
	dl := &Downloader{
		gp: pool,
	}
	return dl, nil
}

type Downloader struct {
	gp *ants.Pool

	// todo: 对出现错误的任务进行缓存, 以便重试
	tasks chan *Task
}

func (dl *Downloader) Stop() {
	dl.gp.Release()
}

func (dl *Downloader) SubmitTask(task Task) error {
	return dl.gp.Submit(task.Func)
}
