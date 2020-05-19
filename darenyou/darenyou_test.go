package darenyou

import (
	"testing"
)

func TestCrawler_PushURL(t *testing.T) {
	c := NewCrawler(Commissioned, SrcO)
	c.PushURL()
}
