package octodex

import (
	"path"
	"testing"

	"github.com/jdxj/wallpaper/models"
)

func TestPath(t *testing.T) {
	name := path.Base("https://octodex.github.com/images/Terracottocat_Single.png")
	t.Log(name)
}

func TestNewOctodexDLI(t *testing.T) {
	cfg := &models.Flags{
		SavePath: "data",
		Retry:    3,
	}
	oDLI := NewOctodexDLI()

	cl := models.NewCrawler(cfg, oDLI)
	cl.Run()
}
