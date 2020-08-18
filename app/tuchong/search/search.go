package search

import (
	"fmt"
	"net/http"

	"github.com/jdxj/wallpaper/app/tuchong"
	"github.com/jdxj/wallpaper/client"

	"github.com/astaxie/beego/logs"
)

const (
	Host = "tuchong.com"

	HomeAPI = "https://tuchong.com/gapi/search/home?%s"
	SiteAPI = "https://tuchong.com/gapi/search/sites?%s"
)

func NewSearch(flags *Flags) *Search {
	s := &Search{
		flags: flags,
		c:     client.New(30),
	}
	return s
}

type Search struct {
	flags *Flags
	c     *http.Client
}

func (s *Search) Query() {
	switch s.flags.Type {
	case Home:
		s.queryHome()
	case User:
		s.querySite()
	}
}

func (s *Search) queryHome() {
	queryURL := fmt.Sprintf(HomeAPI, tuchong.QueryParamsEscape(s.flags.Query))
	httpReq, _ := http.NewRequest(http.MethodGet, queryURL, nil)
	tuchong.SetHTTPReqHeader(httpReq, Host)

	httpResp, err := s.c.Do(httpReq)
	if err != nil {
		logs.Error("%s", err)
		return
	}

	resp, err := UnmarshalResponse(httpResp.Body, true)
	if err != nil {
		logs.Error("%s", err)
		return
	}

	data, err := resp.UnmarshalDataToHome()
	if err != nil {
		logs.Error("%s", err)
		return
	}
	for _, result := range data {
		switch result.Type {
		case Banner:
			entries, err := result.UnmarshalEntriesBanner()
			if err != nil {
				logs.Error("%s", err)
				return
			}
			for _, entry := range entries {
				fmt.Printf("%s\n", entry)
			}
		case Tag:
			entries, err := result.UnmarshalEntriesTag()
			if err != nil {
				logs.Error("%s", err)
				return
			}
			for _, entry := range entries {
				fmt.Printf("%s\n", entry)
			}
		case Site:
			entries, err := result.UnmarshalEntriesSite()
			if err != nil {
				logs.Error("%s", err)
				return
			}
			for _, entry := range entries {
				fmt.Printf("%s\n", entry)
			}
		case Post:
			entries, err := result.UnmarshalEntriesPost()
			if err != nil {
				logs.Error("%s", err)
				return
			}
			for _, entry := range entries {
				fmt.Printf("%s\n", entry)
			}
		case Competition:
			entries, err := result.UnmarshalEntriesCompetition()
			if err != nil {
				logs.Error("%s", err)
				return
			}
			for _, entry := range entries {
				fmt.Printf("%s\n", entry)
			}
		case Course:
			entries, err := result.UnmarshalEntriesCourse()
			if err != nil {
				logs.Error("%s", err)
				return
			}
			for _, entry := range entries {
				fmt.Printf("%s\n", entry)
			}
		}
	}
}

func (s *Search) querySite() {
	queryURL := fmt.Sprintf(SiteAPI, tuchong.QueryParamsEscape(s.flags.Query))
	httpReq, _ := http.NewRequest(http.MethodGet, queryURL, nil)
	tuchong.SetHTTPReqHeader(httpReq, Host)

	httpResp, err := s.c.Do(httpReq)
	if err != nil {
		logs.Error("%s", err)
		return
	}

	resp, err := UnmarshalResponse(httpResp.Body, true)
	if err != nil {
		logs.Error("%s", err)
		return
	}

	sites, err := resp.UnmarshalDataToSite()
	if err != nil {
		logs.Error("%s", err)
		return
	}

	fmt.Printf("%s\n", sites)
}
