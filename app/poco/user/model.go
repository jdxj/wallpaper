package user

import (
	"encoding/json"
	"fmt"
)

func NewParam(id string) *Param {
	p := &Param{
		VisitedUserID: id,
	}
	return p
}

type Param struct {
	UserID        string `json:"user_id"`
	VisitedUserID string `json:"visited_user_id"`
}

func (p *Param) JsonRawMessage() json.RawMessage {
	data, _ := json.Marshal(p)
	return data
}

type User struct {
	UserID        int       `json:"user_id"`
	VisitedUserID int       `json:"visited_user_id"`
	Avatar        string    `json:"avatar"`
	Nickname      string    `json:"nickname"`
	Sex           int       `json:"sex"`
	LocationID    string    `json:"location_id"`
	LoactionName  string    `json:"loaction_name"`
	LocationName  string    `json:"location_name"`
	FollowerCount int       `json:"follower_count"`
	FansCount     int       `json:"fans_count"`
	Signature     string    `json:"signature"`
	Domain        string    `json:"domain"`
	CoverImage    string    `json:"cover_image"`
	IsFollowed    int       `json:"is_followed"`
	Sign          int       `json:"sign"`
	Level         *Level    `json:"level_point_info"`
	Identity      *Identity `json:"identity_info"`
	WorksCount    int       `json:"works_count"`
	MedalCount    int       `json:"medal_count"`
	ShareURL      string    `json:"share_url"`
	Description   string    `json:"description"`
}

func (u *User) Print() {
	fmt.Printf("id: %d, nickname: %s\n",
		u.VisitedUserID, u.Nickname)
}

func UnmarshalUser(data json.RawMessage) (*User, error) {
	user := &User{}
	return user, json.Unmarshal(data, user)
}

type Level struct {
	LevelValue int    `json:"level_value"`
	UserPoints int    `json:"user_points"`
	LevelName  string `json:"level_name"`
}

type Identity struct {
	IsEditor                     int        `json:"is_editor"`
	IsModerator                  int        `json:"is_moderator"`
	ModeratorCategory            int        `json:"moderator_category"`
	ModeratorStr                 string     `json:"moderator_str"`
	PocositeID                   int        `json:"pocosite_id"`
	PocositeName                 int        `json:"pocosite_name"`
	PocositeSubDomain            string     `json:"pocosite_sub_domain"`
	IsPocositeMaster             int        `json:"is_pocosite_master"`
	PocositeMasterStr            string     `json:"pocosite_master_str"`
	IsPocositeRecommendCameraman int        `json:"is_pocosite_recommend_cameraman"`
	IsUserFavourite              int        `json:"is_user_favourite"`
	UserFavouriteCategory        int        `json:"user_favourite_category"`
	UserFavouriteStr             string     `json:"user_favourite_str"`
	CertifyList                  []*Certify `json:"certify_list"`
}

type Certify struct {
	CertifyType  string `json:"certify_type"`
	Category     int    `json:"category"`
	CategoryName string `json:"category_name"`
	CertifyInfo  string `json:"certify_info"`
	Remark       string `json:"remark"`
}
