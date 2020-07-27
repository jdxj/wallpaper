package album

import (
	"fmt"
	"testing"

	"github.com/jdxj/wallpaper/client"

	"github.com/jdxj/wallpaper/models"
)

func TestNewAlbumDLI(t *testing.T) {
	mFlags := &models.CommonFlags{
		SavePath:   "data",
		Retry:      3,
		Concurrent: 3,
	}
	flags := &Flags{
		CommonFlags: mFlags,
		ID:          "5d834a6fe7bce73981fabf4c",
		Limit:       2,
		Adult:       false,
		Order:       New,
	}

	ad := NewAlbumDLI(flags)
	c := client.New()
	ad.SetClient(c)
	var result []string
	if ad.HasNext() {
		t.Log("hash next")
	}

	t.Logf("amount: %d\n", len(result))
	for _, v := range result {
		t.Logf("%s\n", v)
	}
}

func TestAlbumDL_URL(t *testing.T) {
	flags := &Flags{
		ID:    "5d834a6fe7bce73981fabf4c",
		Limit: 2,
		Adult: false,
		Order: New,
	}

	ad := NewAlbumDLI(flags)
	c := client.New()
	ad.SetClient(c)

	dls := ad.Next()
	for _, v := range dls {
		fmt.Printf("%s, %s\n", v.URL(), v.FileName())
	}
}
