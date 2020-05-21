package client

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

func TestGetRemaining(t *testing.T) {
	url := "https://wallhaven.cc/w/96v2dd"
	count := 70
	for i := 0; i < count; i++ {
		if _, err := LimitedGet(url); err != nil {
			fmt.Printf("err: %s\n", err)

		}
	}
}

func TestResetLimiter(t *testing.T) {
	limit := rate.Every(interval)
	fmt.Printf("%f\n", limit)
	fmt.Printf("mid: %v\n", limit*1000+0.5)
	fmt.Printf("%v\n", time.Duration(limit*1000+0.5)*time.Millisecond)
}
