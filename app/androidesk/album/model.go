package album

import (
	"encoding/json"

	"github.com/jdxj/wallpaper/app/androidesk"
)

type Album struct {
	EName   string           `json:"ename"`
	IsFeed  bool             `json:"isfeed"`
	Tags    []string         `json:"tag"`
	ID      string           `json:"id"`
	Top     int              `json:"top"`
	Type    int              `json:"type"`
	Status  string           `json:"status"`
	User    *androidesk.User `json:"user"`
	Favs    int              `json:"favs"`
	Atime   float64          `json:"atime"`
	Desc    string           `json:"desc"`
	Name    string           `json:"name"`
	URLs    []string         `json:"url"`
	Cover   string           `json:"cover"`
	LCover  string           `json:"lcover"`
	SubName string           `json:"subname"`
	SN      int              `json:"sn"`
}

// Wallpapers 是对 android.Wallpaper 的封装
type Wallpapers struct {
	Data []*androidesk.Wallpaper `json:"wallpaper"`
}

func (w *Wallpapers) Unmarshal(data json.RawMessage) error {
	return json.Unmarshal(data, w)
}
