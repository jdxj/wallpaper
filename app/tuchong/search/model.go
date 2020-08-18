package search

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

const (
	Banner      = "banner"
	Tag         = "tag"
	Site        = "site"
	Post        = "post"
	Competition = "competition"
	Course      = "course"
)

var (
	ErrTypeMismatch = errors.New("type mismatch")
)

func UnmarshalResponse(body io.ReadCloser, isGzip bool) (*Response, error) {
	defer body.Close()

	reader := body.(io.Reader)
	if isGzip {
		r, err := gzip.NewReader(reader)
		if err != nil {
			return nil, err
		}
		reader = r
	}

	resp := &Response{}
	decoder := json.NewDecoder(reader)
	return resp, decoder.Decode(resp)
}

type Response struct {
	Result string          `json:"result"`
	Data   json.RawMessage `json:"data"`
}

func (resp *Response) UnmarshalDataToHome() ([]*Result, error) {
	var results []*Result
	return results, json.Unmarshal(resp.Data, &results)
}

func (resp *Response) UnmarshalDataToSite() (*Sites, error) {
	sites := &Sites{}
	return sites, json.Unmarshal(resp.Data, sites)
}

type Sites struct {
	SiteList []*EntrySite `json:"site_list"`
	Sites    int          `json:"sites"` // 当前页的项的数量
	SearchID string       `json:"search_id"`
}

func (s *Sites) String() string {
	buf := bytes.NewBufferString("")
	buf.WriteString(fmt.Sprintf("search_id: %s\n", s.SearchID))

	count := len(s.SiteList)
	for i := 0; i < count; i++ {
		buf.WriteString(fmt.Sprintf("%2d: %s", i, s.SiteList[i]))
		if i != count-1 {
			buf.WriteString("\n")
		}
	}
	return buf.String()
}

type Result struct {
	Type    string          `json:"type"`
	Entries json.RawMessage `json:"entry"`
}

func (res *Result) UnmarshalEntriesBanner() ([]*EntryBanner, error) {
	if res.Type != Banner {
		return nil, ErrTypeMismatch
	}
	var entries []*EntryBanner
	return entries, json.Unmarshal(res.Entries, &entries)
}

func (res *Result) UnmarshalEntriesTag() ([]*EntryTag, error) {
	if res.Type != Tag {
		return nil, ErrTypeMismatch
	}
	var entries []*EntryTag
	return entries, json.Unmarshal(res.Entries, &entries)
}
func (res *Result) UnmarshalEntriesSite() ([]*EntrySite, error) {
	if res.Type != Site {
		return nil, ErrTypeMismatch
	}
	var entries []*EntrySite
	return entries, json.Unmarshal(res.Entries, &entries)
}
func (res *Result) UnmarshalEntriesPost() ([]*EntryPost, error) {
	if res.Type != Post {
		return nil, ErrTypeMismatch
	}
	var entries []*EntryPost
	return entries, json.Unmarshal(res.Entries, &entries)
}
func (res *Result) UnmarshalEntriesCompetition() ([]*EntryCompetition, error) {
	if res.Type != Competition {
		return nil, ErrTypeMismatch
	}
	var entries []*EntryCompetition
	return entries, json.Unmarshal(res.Entries, &entries)
}
func (res *Result) UnmarshalEntriesCourse() ([]*EntryCourse, error) {
	if res.Type != Course {
		return nil, ErrTypeMismatch
	}
	var entries []*EntryCourse
	return entries, json.Unmarshal(res.Entries, &entries)
}

// UnmarshalEntries 仅保留而不应该被使用
func (res *Result) UnmarshalEntries() (interface{}, error) {
	var v interface{}
	switch res.Type {
	case Banner:
		var entries []*EntryBanner
		v = &entries
	case Tag:
		var entries []*EntryTag
		v = &entries
	case Site:
		var entries []*EntrySite
		v = &entries
	case Post:
		var entries []*EntryPost
		v = &entries
	case Competition:
		var entries []*EntryCompetition
		v = &entries
	case Course:
		var entries []*EntryCourse
		v = &entries
	}

	return v, json.Unmarshal(res.Entries, v)
}

// todo: 这种结构可能也有数据
type EntryBanner struct {
}

// todo: 实现
func (eb *EntryBanner) String() string {
	return fmt.Sprintf("EntryBanner don't implement 'String()'")
}

type EntryTag struct {
	TagID          int             `json:"tag_id"`
	TagName        string          `json:"tag_name"`
	Type           string          `json:"type"`
	TagType        string          `json:"tag_type"`
	EventType      string          `json:"event_type"`
	Status         string          `json:"status"`
	Title          string          `json:"title"`
	SubTitle       string          `json:"sub_title"`
	Description    string          `json:"description"`
	CoverImgID     string          `json:"cover_img_id"`
	CoverURL       string          `json:"cover_url"`
	Acl            bool            `json:"acl"`
	AclDesc        string          `json:"acl_desc"`
	CreatedAt      string          `json:"created_at"`
	Subscribers    int             `json:"subscribers"`
	Posts          int             `json:"posts"`
	Participants   int             `json:"participants"`
	Subscribed     bool            `json:"subscribed"`
	Owners         json.RawMessage `json:"owners"`
	ImageUrls      json.RawMessage `json:"image_urls"`
	ApplyStatus    int             `json:"apply_status"`
	TagNameHlp     [][]int         `json:"tag_name_hlp"`
	DescriptionHlp json.RawMessage `json:"description_hlp"`
	PrizeDescHlp   json.RawMessage `json:"prize_desc_hlp"`
}

func (et *EntryTag) String() string {
	return fmt.Sprintf("tag_id: %d, tag_name: %s",
		et.TagID, et.TagName)
}

type EntrySite struct {
	SiteID           string          `json:"site_id"`
	Type             string          `json:"type"`
	Name             string          `json:"name"`
	Description      string          `json:"description"`
	Icon             string          `json:"icon"`
	Verifications    int             `json:"verifications"`
	VerificationList []*Verification `json:"verification_list"`
	Verified         bool            `json:"verified"`
	VerifiedType     int             `json:"verified_type"`
	VerifiedReason   string          `json:"verified_reason"`
	IsFollowing      bool            `json:"is_following"`
	IsFollower       bool            `json:"is_follower"`
	Followers        int             `json:"followers"`
	Appearance       json.RawMessage `json:"appearance"`
	Following        int             `json:"following"`
	Images           json.RawMessage `json:"images"`
	NameHlp          [][]int         `json:"name_hlp"`
	DescriptionHlp   json.RawMessage `json:"description_hlp"`
}

// todo: 实现
func (es *EntrySite) String() string {
	return fmt.Sprintf("site_id: %s, name: %s",
		es.SiteID, es.Name)
}

type Verification struct {
	VerificationType   int    `json:"verification_type"`
	VerificationReason string `json:"verification_reason"`
}

type EntryPost struct {
	PostID      string          `json:"post_id"`
	AuthorID    string          `json:"author_id"`
	Type        string          `json:"type"`
	PublishedAt string          `json:"published_at"`
	Excerpt     string          `json:"excerpt"`
	Favorites   int             `json:"favorites"`
	Comments    int             `json:"comments"`
	Title       string          `json:"title"`
	ImageCount  int             `json:"image_count"`
	Rewardable  bool            `json:"rewardable"`
	Rewards     int             `json:"rewards"`
	Wallpaper   bool            `json:"wallpaper"`
	Views       int             `json:"views"`
	Collected   bool            `json:"collected"`
	Downloads   int             `json:"downloads"`
	Delete      bool            `json:"delete"`
	Update      bool            `json:"update"`
	URL         string          `json:"url"`
	Recommend   bool            `json:"recommend"`
	IsSelf      int             `json:"is_self"`
	Site        *EntrySite      `json:"site"`
	IsFavorite  bool            `json:"is_favorite"`
	Images      []*Image        `json:"images"`
	TitleImage  *TitleImage     `json:"title_image"`
	Content     string          `json:"content"`
	Shares      int             `json:"shares"`
	CollectNum  int             `json:"collect_num"`
	Tags        json.RawMessage `json:"tags"`
	IsTop       bool            `json:"is_top"`
	TitleHlp    [][]int         `json:"title_hlp"`
	ContentHlp  [][]int         `json:"content_hlp"`
	ExcerptHlp  [][]int         `json:"excerpt_hlp"`
	NameHlp     json.RawMessage `json:"name_hlp"`
}

// todo: 实现
func (ep *EntryPost) String() string {
	return fmt.Sprintf("post_id: %s, author_id: %s, title: %s",
		ep.PostID, ep.AuthorID, ep.Title)
}

type Image struct {
	ImgID          int             `json:"img_id"`
	ImgIDStr       string          `json:"img_id_str"`
	UserID         int             `json:"user_id"`
	Title          string          `json:"title"`
	Excerpt        string          `json:"excerpt"`
	Width          int             `json:"width"`
	Height         int             `json:"height"`
	Source         json.RawMessage `json:"source"`
	IsAuthorizedTc int             `json:"is_authorized_tc"`
	IsAuthorTK     bool            `json:"is_author_tk"`
}

type TitleImage struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	ImgID  int    `json:"img_id"`
	URL    string `json:"url"`
}

// todo: 这种结构应该也有数据
type EntryCompetition struct {
}

// todo: 实现
func (ec *EntryCompetition) String() string {
	return fmt.Sprintf("EntryCompetition don't implement 'String()'")
}

// EntryCourse
// 继承 EntryPost 以及 String()
type EntryCourse struct {
	*EntryPost
}
