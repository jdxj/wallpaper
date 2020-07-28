package models

type CommonFlags struct {
	SavePath string
	Retry    int // 失败时的重试次数.
	// 并发数, 即同时下载文件的数量,
	// 该数量也是 http.Client 上每台主机的链接数.
	Concurrent int
	Timeout    int // 下载文件时的超时时间.
}
