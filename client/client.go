package client

import (
	"context"
	"io"
	"net/http"
	"net/http/cookiejar"
	"time"

	"golang.org/x/time/rate"
)

const (
	userAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36"

	interval = 1350 * time.Millisecond
)

var (
	// 全局只使用该 http client
	client  = newClient()
	limiter = rate.NewLimiter(rate.Every(interval), 1)
)

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

// LimitedGet 用于限制请求速度
func LimitedGet(url string) (*http.Response, error) {
	extent := time.Second
	timeout := time.Duration(limiter.Limit()*1000+0.5)*time.Millisecond + extent
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := limiter.Wait(ctx); err != nil {
		return nil, err
	}

	return Get(url)
}

func ResetLimiter(interval time.Duration, cap int) {
	limiter = rate.NewLimiter(rate.Every(interval), cap)
}
