package polayoutu

import (
	"testing"

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
