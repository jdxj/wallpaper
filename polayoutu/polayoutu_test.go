package polayoutu

import (
	"fmt"
	"testing"

	"github.com/jdxj/wallpaper/utils"
)

func TestCrawler_PushURL(t *testing.T) {
	pc := NewCrawler(Thumb)
	go pc.PushURL(182)

	for photo := range pc.urlQueue {
		fmt.Println(photo.Thumb)
	}
}

func TestCrawler_Download(t *testing.T) {
	pc := NewCrawler(Thumb)

	go func() {
		photo := &Photo{
			Thumb: "http://ppe.oss-cn-shenzhen.aliyuncs.com/collections/182/2/thumb.jpg",
		}
		pc.urlQueue <- photo
		close(pc.urlQueue)
	}()

	go pc.Download()

	for photoFile := range pc.photoQueue {
		fmt.Printf("%s\n", photoFile.fileName)
		err := utils.WriteToFileReadCloser(photoFile.fileName, photoFile.data)
		if err != nil {
			t.Fatalf("%s", err)
		}
	}
}

func TestCrawler_WriteToFile(t *testing.T) {
	pc := NewCrawler(Thumb)

	go func() {
		photo := &Photo{
			Thumb: "http://ppe.oss-cn-shenzhen.aliyuncs.com/collections/182/2/thumb.jpg",
		}
		pc.urlQueue <- photo
		close(pc.urlQueue)
	}()

	go pc.Download()

	pc.WriteToFile("data")
}
