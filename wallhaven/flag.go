package wallhaven

import (
	"fmt"
)

type Flags struct {
	//categories, 必有字段
	General bool
	Anime   bool
	People  bool

	//purity, 必有字段
	Sfw     bool
	Sketchy bool
	Nsfw    bool

	// resolution
	// ratio
	// color

	Sorting  string // 必有字段
	TopRange string // 可选字段
	Order    string // 必有字段
	Page     int    // 必有字段

	Path string // 可选
	Url  string // 可选, 指定一个链接进行下载
}

const (
	Path = "data"

	// sorting
	Relevance   = "relevance"
	Random      = "random"
	DateAdded   = "date_added"
	Views       = "views"
	Favorites   = "favorites"
	TopList     = "toplist"
	TopListBeta = "toplist-beta"

	// topRange
	OneDay     = "1d"
	ThreeDay   = "3d"
	OneWeek    = "1w"
	OneMonth   = "1M"
	ThreeMonth = "3M"
	SixMonth   = "6M"
	OneYear    = "1y"

	// order
	Desc = "desc"
	Asc  = "asc"
)

var (
	SortingOptionalValue = fmt.Sprintf("[%s | %s | %s | %s | %s | %s | %s]",
		Relevance, Random, DateAdded, Views, Favorites, TopList, TopListBeta)
	TopRangeOptionalValue = fmt.Sprintf("[%s | %s | %s | %s | %s | %s | %s]",
		OneDay, ThreeDay, OneWeek, OneMonth, ThreeMonth, SixMonth, OneYear)
	OrderOptionalValue = fmt.Sprintf("[%s | %s]", Desc, Asc)
)
