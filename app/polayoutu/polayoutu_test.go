package polayoutu

import (
	"fmt"
	"testing"

	"github.com/jdxj/wallpaper/client"

	"github.com/jdxj/wallpaper/models"
)

func TestNewPoLaYouTuDLI(t *testing.T) {
	flags := &Flags{
		Size:    Thumb,
		Edition: 183,
	}
	mFlags := &models.CommonFlags{
		SavePath: "data",
		Retry:    3,
	}

	pl := NewPoLaYouTuDLI(flags)
	cl := models.NewCrawler(mFlags, pl)
	cl.Run()
}

func TestPlytDL_URL(t *testing.T) {
	flags := &Flags{
		Size:    Thumb,
		Edition: 182,
	}
	pdli := NewPoLaYouTuDLI(flags)
	pdli.c = client.New()

	dls := pdli.Next()
	for _, v := range dls {
		fmt.Printf("%s, %s\n", v.URL(), v.FileName())
	}
}
