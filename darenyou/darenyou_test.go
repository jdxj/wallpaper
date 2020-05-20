package darenyou

import (
	"testing"
)

func TestCrawler_PushURL(t *testing.T) {
	c := NewCrawler(Chaos, SrcO)
	c.PushURL()
}

func TestParseCmd(t *testing.T) {
	params := []string{"-project", Commissioned, "-size", SrcO}
	if err := ParseCmd(params); err != nil {
		t.Fatalf("%s\n", err)
	}

	if project != Commissioned {
		t.Fatalf("project err: %s\n", project)
	}
	if size != SrcO {
		t.Fatalf("size err: %s\n", size)
	}
}
