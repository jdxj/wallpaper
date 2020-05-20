package polayoutu

import (
	"encoding/json"
	"fmt"

	"github.com/jdxj/wallpaper/client"
	"github.com/jdxj/wallpaper/download"
	"github.com/jdxj/wallpaper/utils"
)

const (
	mainPage = "https://www.polayoutu.com/collections/get_entries_by_collection_id/%d?{}"
)

func NewCrawler(cp *CmdParser) *Crawler {
	pc := &Crawler{
		downloader: download.NewDownloader(),
		cmdParser:  cp,
	}
	return pc
}

type Crawler struct {
	downloader *download.Downloader

	cmdParser *CmdParser
}

func (pc *Crawler) PushURL() {
	url := fmt.Sprintf(mainPage, pc.cmdParser.edition)
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
		switch pc.cmdParser.size {
		case FullRes:
			url = photo.FullRes

		case Thumb:
			url = photo.Thumb

		default:
			fmt.Printf("no this size: %s\n", pc.cmdParser.size)
			return
		}

		suffix := utils.TruncateFileName(url)
		fileName := fmt.Sprintf("%d_%d_%d_%s",
			photo.CollectionID, photo.ID, photo.UserID, suffix)

		reqTask := &download.RequestTask{
			Path:     pc.cmdParser.path,
			FileName: fileName,
			URL:      url,
		}
		pc.downloader.PushTask(reqTask)
	}

	pc.downloader.WaitSave()
}
