package login

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/astaxie/beego/logs"
	"github.com/jdxj/wallpaper/db"

	"github.com/jdxj/wallpaper/app/tuchong"
	"github.com/jdxj/wallpaper/client"
)

const (
	Host = "api.tuchong.com"

	TokenAPI = "https://api.tuchong.com/accounts/login?%s"
)

var (
	ErrLoginFail = errors.New("login fail")
)

func NewLogin(flags *Flags) *Account {
	acc := &Account{
		flags: flags,
		c:     client.New(30),
	}
	return acc
}

type Account struct {
	flags *Flags
	c     *http.Client
}

func (acc *Account) RecordNewToken() {
	resp, err := acc.login()
	if err != nil {
		logs.Error("%s", err)
		return
	}

	sqlite := db.Get()
	query := fmt.Sprintf("select count(*) from %s where id=?", tuchong.TableName)
	row := sqlite.QueryRow(query, resp.Identity)

	var count int
	if err := row.Scan(&count); err != nil {
		logs.Error("%s", err)
		return
	}
	if count == 0 { // 增加
		query = fmt.Sprintf("insert into %s (id,token,name) values (?,?,?)", tuchong.TableName)
		_, err := sqlite.Exec(query, resp.Identity, resp.Token, resp.Name)
		if err != nil {
			logs.Error("%s", err)
			return
		}
	} else {
		query = fmt.Sprintf("update %s set token=? where id=?", tuchong.TableName)
		_, err := sqlite.Exec(query, resp.Token, resp.Identity)
		if err != nil {
			logs.Error("%s", err)
			return
		}
	}
	logs.Info("record token success")
}

func (acc *Account) login() (*LoginResponse, error) {
	URL := fmt.Sprintf(TokenAPI, tuchong.QueryParamsEscape(""))
	vs := url.Values{
		"account":  []string{acc.flags.Phone},
		"password": []string{acc.flags.Pass},
	}
	body := strings.NewReader(vs.Encode())

	httpReq, _ := http.NewRequest(http.MethodPost, URL, body)
	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	tuchong.SetHTTPReqHeader(httpReq, Host)

	httpResp, err := acc.c.Do(httpReq)
	if err != nil {
		return nil, err
	}

	resp, err := UnmarshalLoginResponse(httpResp.Body, true)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, ErrLoginFail
	}
	return resp, nil
}
