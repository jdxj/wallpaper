package wallhaven

import (
	"fmt"
	"testing"
	"time"
)

func TestPrintBoolToInt(t *testing.T) {
	fmt.Printf("%d\n", 0)
}

func TestInitialQueryURL(t *testing.T) {
	cp := &CmdParser{
		general:  true,
		anime:    true,
		people:   true,
		sfw:      true,
		sketchy:  true,
		nsfw:     true,
		sorting:  Random,
		topRange: SixMonth,
		order:    Desc,
		page:     1,
	}

	c := &Crawler{
		cmdParser: cp,
		pageURLs:  make(chan string, pageURLLimit),
	}

	c.pageURLs <- "https://wallhaven.cc/w/96v2dd"
	go c.parseURL()

	time.Sleep(3 * time.Second)
	close(c.pageURLs)
}
