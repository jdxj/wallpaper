package client

import (
	"io"
	"net/http"
	"net/http/cookiejar"
)

const userAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36"

// 全局只使用该 http client
var client = newClient()

func newClient() *http.Client {
	jar, _ := cookiejar.New(nil)
	cli := &http.Client{
		Jar: jar,
	}
	return cli
}

func Get(url string) (*http.Response, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("User-Agent", userAgent)

	return client.Do(req)
}

func GetReadCloser(url string) (io.ReadCloser, error) {
	resp, err := Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
