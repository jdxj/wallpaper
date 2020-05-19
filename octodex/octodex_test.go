package octodex

import "testing"

func TestCrawler_PushURL(t *testing.T) {
	c := NewCrawler()
	c.PushURL()
}
