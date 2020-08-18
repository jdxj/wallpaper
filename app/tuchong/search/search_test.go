package search

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/jdxj/wallpaper/app/tuchong"
)

func TestUnmarshalMap(t *testing.T) {
	str := `{"abc":123}`
	//m := make(map[string]interface{})
	var m map[string]interface{}
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
}

func TestUnmarshalResponse(t *testing.T) {
	URL := "https://tuchong.com/gapi/search/home?mac_address=82%3A68%3A67%3A6F%3AC5%3ABC&language=zh&resolution=1080*2208&device_type=Redmi%20K20%20Pro&device_platform=android&os_api=29&device_brand=Xiaomi&openudid=3cefb60f1b2d1769&_rticket=1597729276149&version_code=6132&version_name=6.13.2&ac=wifi&aid=1130&dpi=440&iid=69972369875383&cdid=ab5f288b-d49c-43a8-8be0-91b0ceb25e74&uuid=1130&device_id=3887472264612877&query=%E8%8B%B9%E6%9E%9C&ssmix=a&os_version=10&channel=tengxun&app_name=tuchong&update_version_code=6132&manifest_version_code=6132"
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	tuchong.SetHTTPReqHeader(req, Host)
	c := http.Client{}
	httpResp, err := c.Do(req)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	resp, err := UnmarshalResponse(httpResp.Body, true)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("result: %s\n", resp.Result)

	data, err := resp.UnmarshalDataToHome()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	for _, res := range data {
		switch res.Type {
		case Banner:
			_, err := res.UnmarshalEntriesBanner()
			if err != nil {
				t.Fatalf("%s\n", err)
			}
		case Tag:
			entries, err := res.UnmarshalEntriesTag()
			if err != nil {
				t.Fatalf("%s\n", err)
			}
			for _, entry := range entries {
				fmt.Printf("%#v\n", *entry)
			}
		case Site:
			entries, err := res.UnmarshalEntriesSite()
			if err != nil {
				t.Fatalf("%s\n", err)
			}
			for _, entry := range entries {
				fmt.Printf("%#v\n", *entry)
			}
		case Post:
			entries, err := res.UnmarshalEntriesPost()
			if err != nil {
				t.Fatalf("%s\n", err)
			}
			for _, entry := range entries {
				fmt.Printf("%#v\n", *entry)
			}
		case Competition:
			entries, err := res.UnmarshalEntriesCompetition()
			if err != nil {
				t.Fatalf("%s\n", err)
			}
			for _, entry := range entries {
				fmt.Printf("%#v\n", *entry)
			}
		case Course:
			entries, err := res.UnmarshalEntriesCourse()
			if err != nil {
				t.Fatalf("%s\n", err)
			}
			for _, entry := range entries {
				fmt.Printf("%#v\n", *entry.EntryPost)
			}
		}
	}
}

func TestResult_UnmarshalEntries(t *testing.T) {
	URL := "https://tuchong.com/gapi/search/home?mac_address=82%3A68%3A67%3A6F%3AC5%3ABC&language=zh&resolution=1080*2208&device_type=Redmi%20K20%20Pro&device_platform=android&os_api=29&device_brand=Xiaomi&openudid=3cefb60f1b2d1769&_rticket=1597729276149&version_code=6132&version_name=6.13.2&ac=wifi&aid=1130&dpi=440&iid=69972369875383&cdid=ab5f288b-d49c-43a8-8be0-91b0ceb25e74&uuid=1130&device_id=3887472264612877&query=%E8%8B%B9%E6%9E%9C&ssmix=a&os_version=10&channel=tengxun&app_name=tuchong&update_version_code=6132&manifest_version_code=6132"
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	tuchong.SetHTTPReqHeader(req, Host)
	c := http.Client{}
	httpResp, err := c.Do(req)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	resp, err := UnmarshalResponse(httpResp.Body, true)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("result: %s\n", resp.Result)

	data, err := resp.UnmarshalDataToHome()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	for _, result := range data {
		_, err := result.UnmarshalEntries()
		if err != nil {
			t.Fatalf("%s\n", err)
		}
		switch result.Type {
		case Banner:
		case Tag:
		case Site:
		case Post:
		case Competition:
		case Course:
		}
	}
}

func TestQueryParamsEscape(t *testing.T) {
	res := tuchong.QueryParamsEscape("苹果")
	fmt.Printf("%s\n", res)
}

func TestTimestamp(t *testing.T) {
	now := time.Now()
	fmt.Printf("Unix(): %20d\n", now.Unix())
	fmt.Printf("UnixNano(): %d\n", now.UnixNano())
}

func TestSearch_Query_Home(t *testing.T) {
	flags := &Flags{
		Type:  Home,
		Query: "苹果",
	}
	s := NewSearch(flags)
	s.Query()
}

func TestSearch_Query_Site(t *testing.T) {
	flags := &Flags{
		Type:  User,
		Query: "摄影师郑锐",
	}
	s := NewSearch(flags)
	s.Query()
}
