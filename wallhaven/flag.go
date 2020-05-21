package wallhaven

import (
	"flag"
	"fmt"
)

const (
	flagName        = "wallhaven"
	defaultSavePath = "data"

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
	sortingOptionalValue = fmt.Sprintf("[%s | %s | %s | %s | %s | %s | %s]",
		Relevance, Random, DateAdded, Views, Favorites, TopList, TopListBeta)
	topRangeOptionalValue = fmt.Sprintf("[%s | %s | %s | %s | %s | %s | %s]",
		OneDay, ThreeDay, OneWeek, OneMonth, ThreeMonth, SixMonth, OneYear)
	orderOptionalValue = fmt.Sprintf("[%s | %s]", Desc, Asc)
)

func NewCmdParser() (string, *CmdParser) {
	cp := &CmdParser{}
	return flagName, cp
}

type CmdParser struct {
	//categories, 必有字段
	general bool
	anime   bool
	people  bool

	//purity, 必有字段
	sfw     bool
	sketchy bool
	nsfw    bool

	// resolution
	// ratio
	// color

	sorting  string // 必有字段
	topRange string // 可选字段
	order    string // 必有字段
	page     int    // 必有字段

	path string // 可选
	url  string // 可选, 指定一个链接进行下载
}

func (cp *CmdParser) ParseCmd(params []string) error {
	flagSet := flag.NewFlagSet(flagName, flag.ExitOnError)

	flagSet.BoolVar(&cp.general, "general", false, "set general categories bit.")
	flagSet.BoolVar(&cp.anime, "anime", false, "set anime categories bit.")
	flagSet.BoolVar(&cp.people, "people", false, "set people categories bit.")

	flagSet.BoolVar(&cp.sfw, "sfw", false, "set sfw purity bit.")
	flagSet.BoolVar(&cp.sketchy, "sketchy", false, "set sketchy purity bit")
	flagSet.BoolVar(&cp.nsfw, "nsfw", false, "set nsfw purity bit.")

	flagSet.StringVar(&cp.sorting, "sorting", TopList, "set collation. "+sortingOptionalValue)
	flagSet.StringVar(&cp.topRange, "topRange", "", "set time range. "+topRangeOptionalValue)
	flagSet.StringVar(&cp.order, "order", Desc, "set order. "+orderOptionalValue)
	flagSet.IntVar(&cp.page, "page", 1, "set page number, does not limit the scope.")

	flagSet.StringVar(&cp.path, "path", defaultSavePath, "set storage path.")
	flagSet.StringVar(&cp.url, "url", "", "specify the url of the image to download.")
	return flagSet.Parse(params)
}

func (cp *CmdParser) Run() {
	c := NewCrawler(cp)
	c.PushURL()
}
