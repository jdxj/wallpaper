package album

import "github.com/jdxj/wallpaper/models"

const (
	// order
	New = "new"
	Hot = "hot"
)

type Flags struct {
	*models.CommonFlags
	//List bool // 列出专辑
	ID    string
	Limit int
	Adult bool
	Order string
}
