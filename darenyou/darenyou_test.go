package darenyou

import (
	"testing"

	"github.com/jdxj/wallpaper/models"
)

func TestNewDaRenYouDLI(t *testing.T) {
	flags := &Flags{
		Project: Hysteresis,
		Size:    Src,
	}
	dry := NewDaRenYouDLI(flags)

	mFlags := &models.CommonFlags{
		SavePath: "data",
		Retry:    3,
	}
	cl := models.NewCrawler(mFlags, dry)
	cl.Run()
}
