package wallhaven

import "github.com/jdxj/wallpaper/models"

type Flags struct {
	*models.CommonFlags
	UserName     string
	CollectionID string
	APIKey       string
	Limit        int
}
