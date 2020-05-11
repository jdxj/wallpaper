package polayoutu

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/jdxj/wallpaper/utils"

	"github.com/jdxj/wallpaper/client"
)

func Run() {
	fs := flag.NewFlagSet("polayoutu", flag.ExitOnError)
	kind := fs.String("kind", "thumb", "specified sharpness")
	edition := fs.Int("edition", 1, "get the specified edition")
	if err := fs.Parse(os.Args[2:]); err != nil {
		panic(err)
	}

	//pc := NewCrawler(Thumb)
	pc := NewCrawler(*kind)
	go pc.PushURL(*edition)
	go pc.Download()
	pc.WriteToFile("data")
}

const (
	FullRes = "full_res"
	Thumb   = "thumb"
)

const (
	cacheLimit = 20
	mainPage   = "https://www.polayoutu.com/collections/get_entries_by_collection_id/%d?{}"
)

func NewCrawler(kind string) *Crawler {
	pc := &Crawler{
		cpuNum:     runtime.NumCPU(),
		urlQueue:   make(chan *Photo, cacheLimit),
		photoQueue: make(chan *photoFile, cacheLimit),
		kind:       kind,
	}
	return pc
}

type Crawler struct {
	cpuNum     int
	urlQueue   chan *Photo
	photoQueue chan *photoFile

	kind string
}

func (pc *Crawler) PushURL(edition int) {
	url := fmt.Sprintf(mainPage, edition)
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	rJson := &ResponseJson{}
	if err := decoder.Decode(rJson); err != nil {
		panic(err)
	}

	photos := make([]*Photo, 0)
	if err := json.Unmarshal(rJson.Data, &photos); err != nil {
		panic(err)
	}

	for _, photo := range photos {
		pc.urlQueue <- photo
	}
	close(pc.urlQueue)
}

func (pc *Crawler) Download() {
	wg := &sync.WaitGroup{}
	for i := 0; i < pc.cpuNum; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for photo := range pc.urlQueue {
				pf, err := newPhotoFile(photo, pc.kind)
				if err != nil {
					fmt.Printf("%s\n", err)
					continue
				}

				pc.photoQueue <- pf
			}
		}()
	}
	wg.Wait()
	close(pc.photoQueue)
}

func newPhotoFile(photo *Photo, kind string) (*photoFile, error) {
	var url string
	switch kind {
	case FullRes:
		url = photo.FullRes

	case Thumb:
		url = photo.Thumb

	default:
		return nil, fmt.Errorf("no this kind: %s", kind)
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	index := strings.LastIndex(url, "/")
	suffix := url[index+1:]
	fileName := fmt.Sprintf("%d_%d_%d_%s", photo.CollectionID, photo.ID, photo.UserID, suffix)
	pf := &photoFile{
		property: photo,
		fileName: fileName,
		data:     resp.Body,
	}
	return pf, nil
}

func (pc *Crawler) WriteToFile(path string) {
	// 保存到本地
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		panic(err)
	}

	wg := &sync.WaitGroup{}
	for i := 0; i < pc.cpuNum; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for photoFile := range pc.photoQueue {
				fmt.Printf("writing file: %s\n", photoFile.fileName)

				dir := path + "/" + photoFile.fileName
				err := utils.WriteToFileReadCloser(dir, photoFile.data)
				if err != nil {
					fmt.Printf("%s\n", err)
					continue
				}
			}
		}()
	}
	wg.Wait()
}

type photoFile struct {
	property *Photo
	fileName string
	data     io.ReadCloser
}
