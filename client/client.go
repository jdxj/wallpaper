package client

import (
	"net/http"
	"net/http/cookiejar"
	"time"
)

const (
	UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36"
)

func New(timeout int) *http.Client {
	jar, _ := cookiejar.New(nil)
	c := &http.Client{
		Jar:     jar,
		Timeout: time.Duration(timeout) * time.Second,
	}
	return c
}
