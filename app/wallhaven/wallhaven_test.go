package wallhaven

import (
	"fmt"
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

func abc() (has bool) {
	defer func() {
		fmt.Printf("%t\n", has)
		has = false
	}()

	return true
}

func TestModifyReturnValue(t *testing.T) {
	res := abc()
	fmt.Printf("%t\n", res)
}
