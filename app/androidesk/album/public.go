package album

import "github.com/jdxj/wallpaper/models"

func Run(flags *Flags) {
	ad := NewAlbumDLI(flags)
	cl := models.NewCrawler(flags.CommonFlags, ad)
	cl.Run()
}
