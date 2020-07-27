package wallhaven

import (
	"testing"

	"github.com/jdxj/wallpaper/models"
)

func TestNewWallhavenDLI(t *testing.T) {
	flags := &Flags{
		CommonFlags: &models.CommonFlags{
			SavePath:   "data",
			Retry:      3,
			Concurrent: 1,
		},
		UserName:     "spraayer",
		CollectionID: "482246",
		APIKey:       "lQ69hWxkJb9FWaoQ7zoGkRbMLIIXsVKJ",
		Limit:        1,
	}
	Run(flags)
}
