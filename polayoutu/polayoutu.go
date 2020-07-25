package polayoutu

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/astaxie/beego/logs"
)

const (
	mainPage = "https://www.polayoutu.com/collections/get_entries_by_collection_id/%d?{}"
)

var (
	ErrSizeNotFound = errors.New("size not found")
)

func NewPoLaYouTuDLI(flags *Flags) *PoLaYouTuDLI {
	pl := &PoLaYouTuDLI{
		flags:   flags,
		hasNext: true,
	}
	return pl
}

type PoLaYouTuDLI struct {
	c       *http.Client
	flags   *Flags
	hasNext bool
}

func (pl *PoLaYouTuDLI) SetClient(c *http.Client) {
	pl.c = c
}

func (pl *PoLaYouTuDLI) HasNext() bool {
	return pl.hasNext
}

func (pl *PoLaYouTuDLI) Next() []string {
	dls, err := pl.parseDownloadLinks()
	if err != nil {
		logs.Error("%s", err)
		return nil
	}
	pl.hasNext = false
	return dls
}

func (pl *PoLaYouTuDLI) parseDownloadLinks() ([]string, error) {
	flags := pl.flags
	url := fmt.Sprintf(mainPage, flags.Edition)

	resp, err := pl.c.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	rJson := &ResponseJson{}
	if err := decoder.Decode(rJson); err != nil {
		return nil, err
	}

	photos := make([]*Photo, 0)
	if err := json.Unmarshal(rJson.Data, &photos); err != nil {
		return nil, err
	}
	if len(photos) == 0 {
		logs.Warn("edition may be error")
		return nil, nil
	}

	downloadLinks := make([]string, 0, len(photos))
	for _, photo := range photos {
		var url string
		switch flags.Size {
		case FullRes:
			url = photo.FullRes
		case Thumb:
			url = photo.Thumb
		default:
			return nil, ErrSizeNotFound
		}
		downloadLinks = append(downloadLinks, url)
	}
	return downloadLinks, nil
}
