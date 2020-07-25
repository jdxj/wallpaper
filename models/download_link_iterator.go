package models

import "net/http"

type DownloadLinkIterator interface {
	SetClient(client *http.Client)
	HasNext() bool
	Next() []string // 一次返回一组 download link
}
