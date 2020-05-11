package octodex

import (
	"fmt"
	"testing"
)

func TestNewCrawler(t *testing.T) {
	oc := NewCrawler()
	go oc.PushURL()

	for url := range oc.urlQueue {
		fmt.Println(url)
	}
}

func TestCrawler_Download(t *testing.T) {
	oc := NewCrawler()
	oc.urlQueue <- downloadPrefix + "/images/Terracottocat_Single.png"
	close(oc.urlQueue)

	oc.Download("data")
}
