package octodex

import "github.com/jdxj/wallpaper/models"

func Run(flags *Flags) {
	oDLI := NewOctodexDLI()
	cl := models.NewCrawler(flags.CommonFlags, oDLI)
	cl.Run()
}
