package models

import "net/http"

type DownloadLinkIterator interface {
	SetClient(client *http.Client)
	HasNext() bool
	Next() []DownloadLink // 一次返回一组 download link
}

type DownloadLink interface {
	// 返回下载链接
	URL() string
	// 文件名
	FileName() string
}
