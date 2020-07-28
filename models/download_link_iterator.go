package models

import "net/http"

// DownloadLinkIterator 由于要下载的文件数量可能是未知的,
// 所以需要迭代的方式来不断的获取下载链接.
type DownloadLinkIterator interface {
	// 为了重用 http 链接, 所以为该迭代器设置同一 http.Client
	SetClient(client *http.Client)
	HasNext() bool
	Next() []DownloadLink // 一次返回一组 download link
}

// DownloadLink 由于有些下载链接并不带有文件扩展名,
// 所以这里对下载链接进行抽象, 核心方法是 FileName() 的实现.
type DownloadLink interface {
	// 返回下载链接
	URL() string
	// 文件名可以由文件元信息进行构造, 而文件元信息一般是在获取
	// 下载链接时才能得到, 所以 DownloadLinkIterator.Next()
	// 返回的是 DownloadLink 接口.
	FileName() string
}
