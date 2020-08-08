package poco

import "github.com/jdxj/wallpaper/models"

func Run(flags *Flags) {
	pd := NewPocoDLI(flags)
	cl := models.NewCrawler(flags.CommonFlags, pd)
	cl.Run()
}
