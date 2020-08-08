package poco

import "github.com/jdxj/wallpaper/models"

type Flags struct {
	*models.CommonFlags
	UserID string
	WorkID int
}
