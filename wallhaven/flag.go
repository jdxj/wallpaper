package wallhaven

import (
	"flag"
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

	flagSet.BoolVar(&cp.general, "general", false, "general sign.")
	flagSet.BoolVar(&cp.anime, "anime", false, "anime sign.")
	flagSet.BoolVar(&cp.people, "people", false, "people sign.")

	flagSet.BoolVar(&cp.sfw, "sfw", false, "sfw sign.")
	flagSet.BoolVar(&cp.sketchy, "sketchy", false, "sketchy sign.")
	flagSet.BoolVar(&cp.nsfw, "nsfw", false, "nsfw sign.")

	flagSet.StringVar(&cp.sorting, "sorting", TopList, "todo: sorting usage.")
	flagSet.StringVar(&cp.topRange, "topRange", "", "todo: topRange usage.")
	flagSet.StringVar(&cp.order, "order", Desc, "todo: order usage.")
	flagSet.IntVar(&cp.page, "page", 1, "todo: page usage.")

	flagSet.StringVar(&cp.path, "path", defaultSavePath, "todo: path usage.")
	flagSet.StringVar(&cp.url, "url", "", "todo: url usage.")
	return flagSet.Parse(params)
}

func (cp *CmdParser) Run() {
	c := NewCrawler(cp)
	c.PushURL()
}
