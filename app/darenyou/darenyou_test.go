package darenyou

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/jdxj/wallpaper/models"
)

func TestNewDaRenYouDLI(t *testing.T) {
	flags := &Flags{
		Project: Chaos,
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

func TestDaRenYouDLI_Next(t *testing.T) {
	flags := &Flags{
		Project: Commissioned,
		Size:    SrcO,
	}
	dry := NewDaRenYouDLI(flags)
	dls := dry.Next()
	for _, dl := range dls {
		fmt.Printf("dl: %s\n", dl.URL())
	}
}

func TestFilePath(t *testing.T) {
	res := filepath.Base("https://payload.cargocollective.com/1/14/462928/13412959/DarenYou_13_1100.jpg")
	fmt.Printf("res: %s\n", res)
}

func TestDryDL_FileName(t *testing.T) {
	dd := &dryDL{
		downloadLink: "https://payload.cargocollective.com/1/14/462928/13412959/DarenYou_13_1100.jpg",
	}
	fmt.Printf("%s\n", dd.FileName())
}
