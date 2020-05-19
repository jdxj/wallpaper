package polayoutu

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/jdxj/wallpaper/client"
	"github.com/jdxj/wallpaper/download"
	"github.com/jdxj/wallpaper/utils"
)

func Run() {
	fs := flag.NewFlagSet("polayoutu", flag.ExitOnError)
	kind := fs.String("kind", "thumb", "specified sharpness. you can use [thumb|full_res]")
	edition := fs.Int("edition", 1, "get the specified edition")
	if err := fs.Parse(os.Args[2:]); err != nil {
		panic(err)
	}

	pc := NewCrawler(*kind)
	pc.PushURL(*edition)
}

const (
	FullRes = "full_res"
	Thumb   = "thumb"
)

const (
	mainPage = "https://www.polayoutu.com/collections/get_entries_by_collection_id/%d?{}"
	savePath = "data"
)

func NewCrawler(kind string) *Crawler {
	pc := &Crawler{
		downloader: download.NewDownloader(),
		kind:       kind,
	}
	return pc
}

type Crawler struct {
	downloader *download.Downloader

	kind string
}

func (pc *Crawler) PushURL(edition int) {
	url := fmt.Sprintf(mainPage, edition)
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("PushURL-Get err: %s\n", err)
		return
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	rJson := &ResponseJson{}
	if err := decoder.Decode(rJson); err != nil {
		fmt.Printf("PushURL-Decode err: %s\n", err)
		return
	}

	photos := make([]*Photo, 0)
	if err := json.Unmarshal(rJson.Data, &photos); err != nil {
		fmt.Printf("PushURL-Unmarshal err: %s\n", err)
		return
	}

	for _, photo := range photos {
		var url string
		switch pc.kind {
		case FullRes:
			url = photo.FullRes

		case Thumb:
			url = photo.Thumb

		default:
			fmt.Printf("no this kind: %s\n", pc.kind)
			return
		}

		suffix := utils.TruncateFileName(url)
		fileName := fmt.Sprintf("%d_%d_%d_%s",
			photo.CollectionID, photo.ID, photo.UserID, suffix)

		reqTask := &download.RequestTask{
			Path:     savePath,
			FileName: fileName,
			URL:      url,
		}
		pc.downloader.PushTask(reqTask)
	}

	pc.downloader.WaitSave()
}
