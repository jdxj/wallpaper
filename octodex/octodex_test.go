package octodex

import (
	"path"
	"testing"
)

func TestPath(t *testing.T) {
	name := path.Base("https://octodex.github.com/images/Terracottocat_Single.png")
	t.Log(name)
}

func TestNew(t *testing.T) {
	flags := &Flags{
		Path:  "data",
		Retry: 3,
	}
	oc := New(flags)
	oc.Run()
}
