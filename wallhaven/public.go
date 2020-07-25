package wallhaven

import "github.com/jdxj/wallpaper/models"

func Run(flags *Flags) {
	wd := NewWallhavenDLI(flags)
	cl := models.NewCrawler(&flags.CommonFlags, wd)
	cl.Run()
}
