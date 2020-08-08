package poco

import (
	"fmt"
	"net/http"
	"path"

	"github.com/astaxie/beego/logs"

	"github.com/jdxj/wallpaper/models"
)

const (
	UserAgent      = "Dalvik/2.1.0 (Linux; U; Android 10; Redmi K20 Pro Build/QQ3A.200705.002)"
	Host           = "app-api.poco.cn"
	AcceptEncoding = "gzip"

	WorksAPI = "http://app-api.poco.cn/v1_1/space/get_user_works_list?req=%s"
	WorkAPI  = "http://app-api.poco.cn/v1_1/works/get_works_info?req=%s"
)

func NewPocoDLI(flags *Flags) *PocoDLI {
	pd := &PocoDLI{
		flags:    flags,
		hasWorks: true, // 先获取一次才能知道是否还有
	}
	return pd
}

type PocoDLI struct {
	c     *http.Client
	flags *Flags

	start    int
	hasWorks bool
}

func (pd *PocoDLI) SetClient(c *http.Client) {
	pd.c = c
}

func (pd *PocoDLI) HasNext() bool {
	return pd.hasWorks
}

func (pd *PocoDLI) Next() []models.DownloadLink {
	flags := pd.flags
	if flags.WorkID != 0 {
		dls, err := pd.workDownloadLink(flags.WorkID)
		if err != nil {
			logs.Error("%s", err)
		}
		pd.hasWorks = false
		return dls
	}

	works, err := pd.nextWorks()
	if err != nil {
		logs.Error("%s", err)
		return nil
	}

	result := make([]models.DownloadLink, 0, 30)
	for _, work := range works.List {
		dls, err := pd.workDownloadLink(work.WorksID)
		if err != nil {
			logs.Error("%s", err)
		}
		result = append(result, dls...)
	}
	return result
}

func (pd *PocoDLI) nextWorks() (*Works, error) {
	pw := &ParamWorks{
		Start:         pd.start,
		Length:        15,
		VisitedUserID: pd.flags.UserID,
	}
	req := NewReq(pw.JsonRawMessage())
	query := fmt.Sprintf(WorksAPI, req.Base64Encode())

	httpReq, _ := http.NewRequest(http.MethodGet, query, nil)
	SetHTTPReqHeader(httpReq)

	httpResp, err := pd.c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	resp, err := UnmarshalResponse(httpResp.Body, true)
	if err != nil {
		return nil, err
	}

	works, err := UnmarshalWorks(resp.Data)
	if err != nil {
		return nil, err
	}
	pd.hasWorks = works.HasMore
	pd.start += pw.Length
	return works, nil
}

func (pd *PocoDLI) workDownloadLink(workID int) ([]models.DownloadLink, error) {
	p := &ParamWork{
		UserID:  "",
		WorksID: workID,
	}
	req := NewReq(p.JsonRawMessage())
	query := fmt.Sprintf(WorkAPI, req.Base64Encode())

	httpReq, _ := http.NewRequest(http.MethodGet, query, nil)
	SetHTTPReqHeader(httpReq)

	httpResp, err := pd.c.Do(httpReq)
	if err != nil {
		return nil, err
	}
	resp, err := UnmarshalResponse(httpResp.Body, true)
	if err != nil {
		return nil, err
	}

	wd, err := UnmarshalWorkData(resp.Data)
	if err != nil {
		return nil, err
	}

	photos := wd.WorkInfo.WorksPhotoData
	dls := make([]models.DownloadLink, 0, len(photos))
	for _, photo := range photos {
		downloadLink := fmt.Sprintf("http:%s", photo.MediaInfo.FileURL)
		pl := &pocoDL{
			downloadLink: downloadLink,
		}
		dls = append(dls, pl)
	}
	return dls, nil
}

type pocoDL struct {
	downloadLink string
}

func (pl *pocoDL) URL() string {
	return pl.downloadLink
}

func (pl *pocoDL) FileName() string {
	return path.Base(pl.downloadLink)
}

func SetHTTPReqHeader(req *http.Request) {
	h := req.Header
	h.Set("User-Agent", UserAgent)
	h.Set("Host", Host)
	h.Set("Accept-Encoding", AcceptEncoding)
}
