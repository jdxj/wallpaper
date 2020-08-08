package search

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/logs"

	"github.com/jdxj/wallpaper/app/poco"
)

func NewParamUser(keyword string) *ParamUser {
	p := &ParamUser{
		UserID:  "",
		Start:   0,
		Length:  15,
		Keyword: keyword,
	}
	return p
}

type ParamUser struct {
	UserID  string `json:"user_id"`
	Start   int    `json:"start"`
	Length  int    `json:"length"`
	Keyword string `json:"keyword"`
}

func (pu *ParamUser) JsonRawMessage() json.RawMessage {
	data, _ := json.Marshal(pu)
	return data
}

func NewParamWorks(keyword string) *ParamWorks {
	pu := &ParamUser{
		Length:  24,
		Keyword: keyword,
	}
	pw := &ParamWorks{
		ParamUser: pu,
		WorksType: 0,
	}
	return pw
}

type ParamWorks struct {
	*ParamUser
	WorksType int `json:"works_type"`
}

func (pw *ParamWorks) JsonRawMessage() json.RawMessage {
	data, _ := json.Marshal(pw)
	return data
}

type List struct {
	Data    []json.RawMessage `json:"list"`
	Total   int               `json:"total"`
	HasMore bool              `json:"has_more"`
	UserID  string            `json:"user_id"`
	Keyword string            `json:"keyword"`
}

func (l *List) UnmarshalUsersInfo() ([]*UserInfo, error) {
	result := make([]*UserInfo, 0, len(l.Data))
	for _, d := range l.Data {
		ui := &UserInfo{}
		if err := json.Unmarshal(d, ui); err != nil {
			return nil, err
		}
		result = append(result, ui)
	}
	return result, nil
}

func (l *List) PrintUsersInfo() {
	usersInfo, err := l.UnmarshalUsersInfo()
	if err != nil {
		logs.Error("%s", err)
		return
	}

	for _, user := range usersInfo {
		fmt.Printf("userID: %d, nickname: %s, signature: %s\n",
			user.UserID, user.Nickname, user.Signature)
	}
}

func (l *List) UnmarshalWorksInfo() ([]*WorksInfo, error) {
	result := make([]*WorksInfo, 0, len(l.Data))
	for _, d := range l.Data {
		wi := &WorksInfo{}
		if err := json.Unmarshal(d, wi); err != nil {
			return nil, err
		}
		result = append(result, wi)
	}
	return result, nil
}

func (l *List) PrintWorksInfo() {
	worksInfo, err := l.UnmarshalWorksInfo()
	if err != nil {
		logs.Error("%s", err)
		return
	}

	for _, wi := range worksInfo {
		fmt.Printf("workID: %s, userID: %d, worksType: %d, title: %s\n",
			wi.WorksID, wi.UserID, wi.WorksType, wi.Title)
	}
}

func UnmarshalList(data json.RawMessage) (*List, error) {
	l := &List{}
	return l, json.Unmarshal(data, l)
}

type UserInfo struct {
	UserID              int            `json:"user_id"`
	Mark                int            `json:"mark"`
	Avatar              string         `json:"avatar"`
	Nickname            string         `json:"nickname"`
	Signature           string         `json:"signature"`
	UserRelation        int            `json:"user_relation"`
	CreateTime          int64          `json:"create_time"`
	CreateTimeStr       string         `json:"create_time_str"`
	Identity            *poco.Identity `json:"user_identity_info"`
	VisitorFollowStatus int            `json:"visitor_follow_status"`
	IsFollowed          int            `json:"is_followed"`
}

type WorksInfo struct {
	*poco.Work
	WorksID    string            `json:"works_id"`
	ImageCount int               `json:"image_count"`
	Tag        []json.RawMessage `json:"tag"`
	Mark       int               `json:"mark"`
}
