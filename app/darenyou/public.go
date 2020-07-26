package darenyou

import "github.com/jdxj/wallpaper/models"

func Run(flags *Flags) {
	dry := NewDaRenYouDLI(flags)
	cl := models.NewCrawler(flags.CommonFlags, dry)
	cl.Run()
}
