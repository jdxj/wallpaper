package octodex

import "testing"

func TestCrawler_PushURL(t *testing.T) {
	flags := &Flags{
		Path: Path,
	}
	c := NewCrawler(flags)
	c.PushURL()
}

func TestPushURLFromWeb(t *testing.T) {
	c := NewCrawler(nil)
	c.pushURLFromWeb()
}
