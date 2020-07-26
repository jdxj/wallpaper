package polayoutu

import "github.com/jdxj/wallpaper/models"

func Run(flags *Flags) {
	pl := NewPoLaYouTuDLI(flags)
	cl := models.NewCrawler(flags.CommonFlags, pl)
	cl.Run()
}
