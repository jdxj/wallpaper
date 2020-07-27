package octodex

import (
	"fmt"
	"path"
	"testing"

	"github.com/jdxj/wallpaper/client"

	"github.com/jdxj/wallpaper/models"
)

func TestPath(t *testing.T) {
	name := path.Base("https://octodex.github.com/images/Terracottocat_Single.png")
	t.Log(name)
}

func TestNewOctodexDLI(t *testing.T) {
	cfg := &models.CommonFlags{
		SavePath: "data",
		Retry:    3,
	}
	oDLI := NewOctodexDLI()

	cl := models.NewCrawler(cfg, oDLI)
	cl.Run()
}

func TestOctodexDL_URL(t *testing.T) {
	oDLI := NewOctodexDLI()
	oDLI.c = client.New()
	for _, v := range oDLI.Next() {
		fmt.Printf("%s, %s\n", v.URL(), v.FileName())
	}
}
