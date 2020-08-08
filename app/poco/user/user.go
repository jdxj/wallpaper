package user

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/logs"

	"github.com/jdxj/wallpaper/app/poco"

	"github.com/jdxj/wallpaper/client"
)

const (
	IdentityAPI = "http://app-api.poco.cn/v1_1/space/get_user_identity_info?req=%s"
)

func NewInfo(flags *Flags) *Info {
	c := client.New(30)
	info := &Info{
		c:     c,
		flags: flags,
	}
	return info
}

type Info struct {
	c     *http.Client
	flags *Flags
}

func (info *Info) Query() {
	flags := info.flags

	if flags.Identity {
		info.queryIdentity()
	}
}

func (info *Info) queryIdentity() {
	p := NewParam(info.flags.ID)
	req := poco.NewReq(p.JsonRawMessage())
	query := fmt.Sprintf(IdentityAPI, req.Base64Encode())

	httpReq, _ := http.NewRequest(http.MethodGet, query, nil)
	poco.SetHTTPReqHeader(httpReq)

	c := info.c
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
	user, err := UnmarshalUser(resp.Data)
	if err != nil {
		logs.Error("%s", err)
		return
	}
	user.Print()
}
