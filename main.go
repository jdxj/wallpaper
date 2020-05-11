package main

import (
	"fmt"
	"time"

	"github.com/jdxj/wallpaper/octodex"
)

func main() {
	start := time.Now()
	oc := octodex.NewCrawler()
	go oc.PushURL()
	go oc.Download()

	oc.Write("data2")
	end := time.Now()

	fmt.Printf("start: %s\n", start)
	fmt.Printf("end: %s\n", end)
	fmt.Printf("expend: %s", end.Sub(start))

}
