package darenyou

import (
	"testing"

	"github.com/jdxj/wallpaper/models"
)

func TestNewDaRenYouDLI(t *testing.T) {
	flags := &Flags{
		Project: Chaos,
		Size:    Src,
	}
	dry := NewDaRenYouDLI(flags)

	mFlags := &models.Flags{
		SavePath: "data",
		Retry:    3,
	}
	cl := models.NewCrawler(mFlags, dry)
	cl.Run()
}
