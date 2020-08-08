package search

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/logs"
	"github.com/jdxj/wallpaper/app/poco"

	"github.com/jdxj/wallpaper/client"
)

const (
	SearchUserAPI  = "http://app-api.poco.cn/v1_1/search/get_user_list?req=%s"
	SearchWorksAPI = "http://app-api.poco.cn/v1_1/search/get_works_list?req=%s"
)

func NewSearch(flags *Flags) *Search {
	c := client.New(30)
	s := &Search{
		c:     c,
		flags: flags,
	}
	return s
}

type Search struct {
	c     *http.Client
	flags *Flags
}

func (sea *Search) Query() {
	flags := sea.flags
	switch flags.Type {
	case User:
		sea.queryUser()
	case Works:
		sea.queryWorks()
	default:
		logs.Warn("no such type: %s", flags.Type)
	}
}

func (sea *Search) queryUser() {
	p := NewParamUser(sea.flags.Keyword)
	req := poco.NewReq(p.JsonRawMessage())
	query := fmt.Sprintf(SearchUserAPI, req.Base64Encode())

	httpReq, _ := http.NewRequest(http.MethodGet, query, nil)
	poco.SetHTTPReqHeader(httpReq)

	c := sea.c
	httpResp, err := c.Do(httpReq)
	if err != nil {
		logs.Error("%s", err)
		return
	}

	resp, err := poco.UnmarshalResponse(httpResp.Body, true)
	if err != nil {
		logs.Error("%s", err)
		return
	}
	list, err := UnmarshalList(resp.Data)
	if err != nil {
		logs.Error("%s", err)
		return
	}
	list.PrintUsersInfo()
}

func (sea *Search) queryWorks() {
	p := NewParamWorks(sea.flags.Keyword)
	req := poco.NewReq(p.JsonRawMessage())
	query := fmt.Sprintf(SearchWorksAPI, req.Base64Encode())

	httpReq, _ := http.NewRequest(http.MethodGet, query, nil)
	poco.SetHTTPReqHeader(httpReq)

	c := sea.c
	httpResp, err := c.Do(httpReq)
	if err != nil {
		logs.Error("%s", err)
		return
	}

	resp, err := poco.UnmarshalResponse(httpResp.Body, true)
	if err != nil {
		logs.Error("%s", err)
		return
	}
	list, err := UnmarshalList(resp.Data)
	if err != nil {
		logs.Error("%s", err)
		return
	}
	list.PrintWorksInfo()
}
