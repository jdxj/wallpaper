package poco

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jdxj/wallpaper/client"
)

func TestJson(t *testing.T) {
	data, err := json.Marshal("喜之狼")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", data)
}

func TestReq_MarshalReq(t *testing.T) {
	param := `{"user_id":"","start":0,"length":15,"keyword":"11111111111111111111"}`
	req := NewReq(json.RawMessage(param))
	fmt.Printf("sign_code: %s\n", req.SignCode)
	res := req.Base64Encode()
	fmt.Printf("%s\n", res)
}

func TestNextWorks(t *testing.T) {
	flags := &Flags{
		CommonFlags: nil,
		UserID:      "179203067",
		WorkID:      0,
	}
	pd := NewPocoDLI(flags)
	pd.SetClient(client.New(30))
	works, err := pd.nextWorks()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	for _, work := range works.List {
		fmt.Printf("worksid: %d, titile: %s\n",
			work.WorksID, work.Title)
	}
}

func TestWorkDownloadLink(t *testing.T) {
	flags := &Flags{
		CommonFlags: nil,
		UserID:      "179203067",
		WorkID:      0,
	}
	pd := NewPocoDLI(flags)
	pd.SetClient(client.New(30))
	works, err := pd.nextWorks()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	for i, work := range works.List {
		fmt.Printf("i: %d, worksid: %d, titile: %s----------------\n",
			i, work.WorksID, work.Title)

		dls, err := pd.workDownloadLink(work.WorksID)
		if err != nil {
			t.Logf("err: %s\n", err)
			continue
		}

		for _, dl := range dls {
			fmt.Printf("fileName: %s, downloadLink: %s\n",
				dl.FileName(), dl.URL())
		}
	}
}

func TestPocoDLI_Next(t *testing.T) {
	flags := &Flags{
		CommonFlags: nil,
		UserID:      "174200527",
		WorkID:      0,
	}
	pd := NewPocoDLI(flags)
	pd.SetClient(client.New(30))

	for i := 0; pd.HasNext(); i++ {
		fmt.Printf("i: %d-------------------------\n", i)
		dls := pd.Next()
		for j, dl := range dls {
			fmt.Printf("j: %d, fileName: %s, downloadLink: %s\n",
				j, dl.FileName(), dl.URL())
		}
	}
}

func TestWorkDownloadLinkArticle(t *testing.T) {
	pd := NewPocoDLI(nil)
	pd.SetClient(client.New(30))
	pd.workDownloadLink(12226919)
}
