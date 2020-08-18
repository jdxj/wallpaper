package tuchong

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/jdxj/wallpaper/models"

	"github.com/jdxj/wallpaper/db"

	"github.com/astaxie/beego/logs"
)

const (
	Host = "api.tuchong.com"

	TableName = "tuchong_account"

	FollowAPI = "https://api.tuchong.com/users/self/following/%s?%s"
	WorksAPI  = "https://api.tuchong.com/sites/%s/works?%s"
)

var (
	SortedQueryParamKes = []string{
		"mac_address",
		"language",
		"resolution",
		"device_type",
		"device_platform",
		"os_api",
		"device_brand",
		"openudid",
		"_rticket",
		"version_code",
		"version_name",
		"ac",
		"aid",
		"dpi",
		"iid",
		"cdid",
		"page", //
		"uuid",
		"device_id",
		"query", //
		"ssmix",
		"before_timestamp", //
		"os_version",
		"channel",
		"app_name",
		"update_version_code",
		"manifest_version_code",
	}

	QueryParams = map[string]string{
		"mac_address":           "82%3A68%3A67%3A6F%3AC5%3ABC",
		"language":              "zh",
		"resolution":            "1080*2208",
		"device_type":           "Redmi%20K20%20Pro",
		"device_platform":       "android",
		"os_api":                "29",
		"device_brand":          "Xiaomi",
		"openudid":              "3cefb60f1b2d1769",
		"_rticket":              "", // 毫秒级时间戳, 需要动态设置
		"version_code":          "6132",
		"version_name":          "6.13.2",
		"ac":                    "wifi",
		"aid":                   "1130",
		"dpi":                   "440",
		"iid":                   "69972369875383",
		"cdid":                  "ab5f288b-d49c-43a8-8be0-91b0ceb25e74",
		"page":                  "1", //
		"uuid":                  "1130",
		"device_id":             "3887472264612877",
		"query":                 "", // 要查询的信息
		"ssmix":                 "a",
		"before_timestamp":      "0", //
		"os_version":            "10",
		"channel":               "tengxun",
		"app_name":              "tuchong",
		"update_version_code":   "6132",
		"manifest_version_code": "6132",
	}

	ErrNotFoundToken = errors.New("not found token")
	ErrFollowFail    = errors.New("follow fail")
)

const (
	//Host           = "tuchong.com"
	Device         = "3887472264612877"
	Version        = "6132"
	Channel        = "tengxun"
	Platform       = "android"
	AcceptEncoding = "gzip"
	UserAgent      = "okhttp/3.12.2"
)

func SetHTTPReqHeader(req *http.Request, host string) {
	h := req.Header
	h.Set("Host", host)
	h.Set("device", Device)
	h.Set("version", Version)
	h.Set("channel", Channel)
	h.Set("platform", Platform)
	h.Set("Accept-Encoding", AcceptEncoding)
	h.Set("User-Agent", UserAgent)
	h.Set("host-name", host)

	// HostAddress 可设置也可不设置,
	// 如果设置, 其值应为 "tuchong.com" 的实际 ip.
	hostAddress, err := getAHostAddress(host)
	if err != nil {
		logs.Error("%s", err)
	} else {
		h.Set("host-address", hostAddress)
	}
}

func getAHostAddress(host string) (string, error) {
	addrs, err := net.LookupHost(host)
	if err != nil {
		return "", err
	}
	return addrs[0], nil
}

// QueryParamsEscape 如果不进行查询, query 应为空串.
func QueryParamsEscape(query string) string {
	rticket := time.Now().UnixNano() / 1000000
	QueryParams["_rticket"] = strconv.Itoa(int(rticket))
	QueryParams["query"] = url.QueryEscape(query)

	// todo: 是否需要一个 string pool?
	buf := bytes.NewBufferString("")
	count := len(SortedQueryParamKes)
	for i := 0; i < count; i++ {
		k := SortedQueryParamKes[i]
		v := QueryParams[k]

		buf.WriteString(k)
		buf.WriteString("=")
		buf.WriteString(v)

		if i != count-1 {
			buf.WriteString("&")
		}
	}
	return buf.String()
}

func NewTuChongDLI(flags *Flags) *TuChongDLI {
	tcd := &TuChongDLI{
		flags: flags,
	}
	return tcd
}

type TuChongDLI struct {
	flags *Flags
	c     *http.Client

	hasFollow bool
	hasNext   bool
	page      int

	beforeTimestamp string
}

func (tcd *TuChongDLI) follow() error {
	URL := fmt.Sprintf(FollowAPI, tcd.flags.SiteID, QueryParamsEscape(""))
	httpReq, _ := http.NewRequest(http.MethodPut, URL, nil)
	token, err := tcd.getToken()
	if err != nil {
		return err
	}
	httpReq.Header.Set("token", token)
	SetHTTPReqHeader(httpReq, Host)

	httpResp, err := tcd.c.Do(httpReq)
	if err != nil {
		return err
	}
	resp, err := UnmarshalFollowResponse(httpResp.Body, false)
	if err != nil {
		return err
	}
	if resp.Result != "SUCCESS" {
		return ErrFollowFail
	}
	return nil
}

func (tcd *TuChongDLI) getToken() (string, error) {
	sqlite := db.Get()
	query := fmt.Sprintf("select token from %s limit 1", TableName)
	row := sqlite.QueryRow(query)
	var token string
	if err := row.Scan(&token); err != nil {
		return "", err
	}
	if token == "" {
		return "", ErrNotFoundToken
	}
	return token, nil
}

func (tcd *TuChongDLI) SetClient(c *http.Client) {
	tcd.c = c
}

func (tcd *TuChongDLI) HasNext() bool {

	return true
}

func (tcd *TuChongDLI) Next() []models.DownloadLink {
	return nil
}
