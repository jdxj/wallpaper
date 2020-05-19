package polayoutu

import "testing"

func TestCrawler_PushURL(t *testing.T) {
	c := NewCrawler(Thumb)
	c.PushURL(182)
}
