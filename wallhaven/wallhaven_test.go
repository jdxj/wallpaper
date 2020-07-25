package wallhaven

import (
	"fmt"
	"testing"
)

func TestPrintBoolToInt(t *testing.T) {
	fmt.Printf("%d\n", 0)
}

func TestInitialQueryURL(t *testing.T) {
	flags := &Flags{
		General:  true,
		Anime:    true,
		People:   true,
		Sfw:      true,
		Sketchy:  true,
		Nsfw:     true,
		Sorting:  Random,
		TopRange: SixMonth,
		Order:    Desc,
		Page:     1,
	}

	c := &WallhavenDLI{
		flags: flags,
	}

	result := c.initialQueryURL()
	fmt.Printf("%s\n", result)
}
