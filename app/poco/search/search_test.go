package search

import (
	"fmt"
	"testing"
)

func TestSearch_Query(t *testing.T) {
	flags := &Flags{
		Type:    User,
		Keyword: "方托马斯",
	}
	sea := NewSearch(flags)
	sea.Query()
}

func TestParamWorks_JsonRawMessage(t *testing.T) {
	pw := NewParamWorks("孤独的一棵树")
	fmt.Printf("%s\n", pw.JsonRawMessage())
}
