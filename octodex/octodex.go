package octodex

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/jdxj/wallpaper/client"
	"github.com/jdxj/wallpaper/utils"
)

func Run() {
	start := time.Now()
	oc := NewCrawler()
	go oc.PushURL()
	go oc.Download()

	oc.Write("data")
	end := time.Now()

	fmt.Printf("start:  %s\n", start)
	fmt.Printf("end:    %s\n", end)
	fmt.Printf("expend: %s", end.Sub(start))
}

const mainPage = "https://octodex.github.com"
const downloadPrefix = mainPage
const cacheLimit = 100

func NewCrawler() *Crawler {
	return &Crawler{
		cpuCount:  runtime.NumCPU(),
		urlQueue:  make(chan string, cacheLimit),
		dataQueue: make(chan *myFile, cacheLimit),
	}
}

type Crawler struct {
	cpuCount int

	urlQueue  chan string
	dataQueue chan *myFile
}

// PushURL 不断地获取下载链接
func (oc *Crawler) PushURL() {
	fmt.Println("start")

	resp, err := client.Get(mainPage)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	sel := doc.Find(".width-fit")
	sel.Each(func(i int, selI *goquery.Selection) {
		src, ok := selI.Attr("data-src")
		if !ok {
			return
		}

		fmt.Printf("pushing     [%s]\n", src)
		oc.urlQueue <- downloadPrefix + src
	})

	close(oc.urlQueue)
}

func (oc *Crawler) Download() {
	// 多协程下载
	wg := &sync.WaitGroup{}
	for i := 0; i < oc.cpuCount; i++ {
		wg.Add(1)

		// 访问远程
		go func() {
			defer wg.Done()

			for url := range oc.urlQueue {
				fmt.Printf("downloading [%s]\n", url)

				myFile, err := newMyFile(url)
				if err != nil {
					fmt.Printf("%s\n", err)
					continue
				}
				oc.dataQueue <- myFile
			}
		}()
	}
	wg.Wait()
	close(oc.dataQueue)
}

func (oc *Crawler) Write(path string) {
	// 保存到本地
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		panic(err)
	}

	// 多协程写
	wg := &sync.WaitGroup{}
	for i := 0; i < oc.cpuCount; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for myFile := range oc.dataQueue {
				fmt.Printf("writing     [%s]\n", myFile.fileName)

				err := utils.WriteToFile(path+"/"+myFile.fileName, myFile.data)
				if err != nil {
					fmt.Println(err)
					continue
				}
			}
		}()
	}
	wg.Wait()
	fmt.Println("all write finish")
}

type myFile struct {
	fileName string
	data     []byte
}

func newMyFile(url string) (*myFile, error) {
	data, err := utils.Download(url)
	if err != nil {
		return nil, err
	}

	// 获取图片名
	start := strings.LastIndex(url, "/")
	if start < 0 {
		return nil, fmt.Errorf("can not find file name")
	}
	fileName := url[start+1:]

	myFile := &myFile{
		fileName: fileName,
		data:     data,
	}
	return myFile, nil
}
