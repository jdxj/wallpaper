package polayoutu

import "github.com/jdxj/wallpaper/models"

const (
	FullRes = "full"
	Thumb   = "thumb"

	EditionNum = 1
)

type Flags struct {
	*models.CommonFlags
	Size    string
	Edition int
}
