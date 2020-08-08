package poco

import (
	"compress/gzip"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
)

func NewReq(param json.RawMessage) *Request {
	req := &Request{
		Param:   param,
		Version: "2.4.4",
		CTime:   1,
		IsEnc:   0,
		AppName: "poco_photography_android",
		OSType:  "android",
	}
	req.signCode()
	return req
}

type Request struct {
	SignCode string          `json:"sign_code"`
	Param    json.RawMessage `json:"param"`
	Version  string          `json:"version"`
	CTime    int64           `json:"ctime"`
	IsEnc    int             `json:"is_enc"`
	AppName  string          `json:"app_name"`
	OSType   string          `json:"os_type"`
}

func (req *Request) Base64Encode() string {
	data := req.marshalReq()
	return base64.StdEncoding.EncodeToString(data)
}

func (req *Request) marshalReq() []byte {
	data, _ := json.Marshal(req)
	return data
}

func (req *Request) signCode() {
	signData := fmt.Sprintf("poco_%s_app", req.Param)
	sign := md5.Sum([]byte(signData))
	signStr := fmt.Sprintf("%x", sign)
	req.SignCode = signStr[5:24]
}

type Response struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
	Runtime string          `json:"runtime"`
	Base    string          `json:"base"`
	AppVer  string          `json:"appver"`
	ApiVer  string          `json:"apiver"`
	Pass    int             `json:"pass"`
}

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

type Works struct {
	Total   int      `json:"total"`
	HasMore bool     `json:"has_more"`
	List    []*Work  `json:"list"`
	Filter  []string `json:"filter"`
}

func UnmarshalWorks(data json.RawMessage) (*Works, error) {
	ws := &Works{}
	return ws, json.Unmarshal(data, ws)
}

type Work struct {
	WorksID              int       `json:"works_id"`
	UserID               int       `json:"user_id"`
	LikeCount            int       `json:"like_count"`
	ClickCount           int       `json:"click_count"`
	CommentCount         int       `json:"comment_count"`
	CreateTime           int64     `json:"create_time"`
	PrivacyStatus        int       `json:"privacy_status"`
	IsOld                int       `json:"is_old"`
	UserNickname         string    `json:"user_nickname"`
	UserAvatar           string    `json:"user_avatar"`
	UserSignature        string    `json:"user_signature"`
	UserFavouriteSign    int       `json:"user_favourite_sign"`
	UserFamousSign       int       `json:"user_famous_sign"`
	UserLevelName        string    `json:"user_level_name"`
	UserIsModerator      int       `json:"user_is_moderator"`
	UserIsPocositeMaster int       `json:"user_is_pocosite_master"`
	UserIdentityInfo     *Identity `json:"user_identity_info"`
	CameraBrandName      string    `json:"camera_brand_name"`
	CameraModelName      string    `json:"camera_model_name"`
	WorksType            int       `json:"works_type"`
	Title                string    `json:"title"`
	Description          string    `json:"description"`
	TagNames             string    `json:"tag_names"`
	CoverImage           int       `json:"cover_image"`
	CoverImageClip       string    `json:"cover_image_clip"`
	CoverImageInfo       *Cover    `json:"cover_image_info"`
	WorksPhotoCount      int       `json:"works_photo_count"`
	IsMedal              int       `json:"is_medal"`
	IsEssence            int       `json:"is_essence"`
	WorksEssenceInfo     *Essence  `json:"works_essence_info"`
	CreateTimeStr        string    `json:"create_time_str"`
	VisitorLikeStatus    int       `json:"visitor_like_status"`
	VisitorCollectStatus int       `json:"visitor_collect_status"`
	VisitorFollowStatus  int       `json:"visitor_follow_status"`
	WorksURL             string    `json:"works_url"`
	IsAudit              int       `json:"is_audit"`
	IsTop                int       `json:"is_top"`
}

type Identity struct {
	IsModerator           int        `json:"is_moderator"`
	ModeratorCategory     int        `json:"moderator_category"`
	ModeratorStr          string     `json:"moderator_str"`
	IsEditor              int        `json:"is_editor"`
	IsPocositeMaster      int        `json:"is_pocosite_master"`
	PocositeID            int        `json:"pocosite_id"`
	PocositeMasterStr     string     `json:"pocosite_master_str"`
	IsUserFavourite       int        `json:"is_user_favourite"`
	UserFavouriteCategory int        `json:"user_favourite_category"`
	UserFavouriteStr      string     `json:"user_favourite_str"`
	CertifyList           []*Certify `json:"certify_list"`
}

type Cover struct {
	MediaID      int    `json:"media_id"`
	UserID       int    `json:"user_id"`
	FileType     int    `json:"file_type"`
	FileURL      string `json:"file_url"`
	FileName     string `json:"file_name"`
	FileExt      string `json:"file_ext"`
	FileSize     int    `json:"file_size"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	DeletedImage string `json:"deleted_image"`
}

type Essence struct {
	CreateTime    int64  `json:"create_time"`
	CreateTimeStr string `json:"create_time_str"`
}

type Certify struct {
	CertifyType  string `json:"certify_type"`
	Category     int    `json:"category"`
	CategoryName string `json:"category_name"`
	CertifyInfo  string `json:"certify_info"`
	Remark       string `json:"remark"`
}

type ParamWorks struct {
	UserID        string `json:"user_id"`
	Start         int    `json:"start"`
	Length        int    `json:"length"`
	WorksType     int    `json:"works_type"`
	VisitedUserID string `json:"visited_user_id"`
}

func (pw *ParamWorks) JsonRawMessage() json.RawMessage {
	data, _ := json.Marshal(pw)
	return data
}

type ParamWork struct {
	UserID  string `json:"user_id"`
	WorksID int    `json:"works_id"`
}

func (pk *ParamWork) JsonRawMessage() json.RawMessage {
	data, _ := json.Marshal(pk)
	return data
}

type WorkData struct {
	WorkInfo  *WorkInfo `json:"works_info"`
	NextWorks string    `json:"next_works"`
}

func UnmarshalWorkData(data json.RawMessage) (*WorkData, error) {
	wd := &WorkData{}
	return wd, json.Unmarshal(data, wd)
}

type WorkInfo struct {
	WorksID              int       `json:"works_id"`
	UserID               int       `json:"user_id"`
	Category             int       `json:"category"`
	RecommendCount       int       `json:"recommend_count"`
	CommentCount         int       `json:"comment_count"`
	LikeCount            int       `json:"like_count"`
	ShareCount           int       `json:"share_count"`
	ViewCount            int       `json:"view_count"`
	ClickCount           int       `json:"click_count"`
	CollectCount         int       `json:"collect_count"`
	CameraBrandID        int       `json:"camera_brand_id"`
	CameraModelID        int       `json:"camera_model_id"`
	WorksStatus          int       `json:"works_status"`
	CreateTime           int64     `json:"create_time"`
	PrivacyStatus        int       `json:"privacy_status"`
	IsOld                int       `json:"is_old"`
	IsHomepageRecommend  int       `json:"is_homepage_recommend"`
	IsEditorRecommend    int       `json:"is_editor_recommend"`
	IsRegionRecommend    int       `json:"is_region_recommend"`
	IsModeratorRecommend int       `json:"is_moderator_recommend"`
	UserNickname         string    `json:"user_nickname"`
	UserAvatar           string    `json:"user_avatar"`
	UserSignature        string    `json:"user_signature"`
	UserFavouriteSign    int       `json:"user_favourite_sign"`
	UserFamousSign       int       `json:"user_famous_sign"`
	UserLevelName        string    `json:"user_level_name"`
	UserIsModerator      int       `json:"user_is_moderator"`
	UserIsPocositeMaster int       `json:"user_is_pocosite_master"`
	UserIdentityInfo     *Identity `json:"user_identity_info"`
	CameraBrandName      string    `json:"camera_brand_name"`
	CameraModelName      string    `json:"camera_model_name"`
	CreateTimeStr        string    `json:"create_time_str"`
	Title                string    `json:"title"`
	Description          string    `json:"description"`
	Tag                  string    `json:"tag"`
	Copyright            int       `json:"copyright"`
	CoverImage           int       `json:"cover_image"`
	CoverImageClip       string    `json:"cover_image_clip"`
	BgMusicID            int       `json:"bg_music_id"`
	BgMusic              string    `json:"bg_music"`
	ImageSort            string    `json:"image_sort"`
	WorksType            int       `json:"works_type"`
	CreateSource         int       `json:"create_source"`
	ShootLocation        string    `json:"shoot_location"`
	PoiType              string    `json:"poi_type"`
	DetailLocation       string    `json:"detail_location"`
	Latitude             string    `json:"latitude"`
	Longitude            string    `json:"longitude"`
	WorksHeatElse        string    `json:"works_heat_else"`
	IsReprint            int       `json:"is_reprint"`
	ReprintURL           string    `json:"reprint_url"`
	CategoryName         string    `json:"category_name"`
	CopyrightStr         string    `json:"copyright_str"`
	TagNames             string    `json:"tag_names"`
	CoverImageInfo       *Cover    `json:"cover_image_info"`
	WorksPhotoCount      int       `json:"works_photo_count"`
	IsMedal              int       `json:"is_medal"`
	WorksMedalInfo       []*Medal  `json:"works_medal_info"`
	IsEssence            int       `json:"is_essence"`
	WorksEssenceInfo     *Essence  `json:"works_essence_info"`
	WorksPhotoData       []*Photo  `json:"works_photo_data"`
	WorksCommentAccess   int       `json:"works_comment_access"`
	VisitorLikeStatus    int       `json:"visitor_like_status"`
	VisitorCollectStatus int       `json:"visitor_collect_status"`
	VisitorFollowStatus  int       `json:"visitor_follow_status"`
	WorksURL             string    `json:"works_url"`
}

type Medal struct {
	WorksID       int    `json:"works_id"`
	MedalID       int    `json:"medal_id"`
	Count         int    `json:"count"`
	CreateTime    int64  `json:"create_time"`
	MedalName     string `json:"medal_name"`
	MedalIcon     string `json:"medal_icon"`
	CreateTimeStr string `json:"create_time_str"`
}

type Photo struct {
	MediaID            int    `json:"media_id"`
	WorksID            int    `json:"works_id"`
	Description        string `json:"description"`
	FavCount           int    `json:"fav_count"`
	ViewCount          int    `json:"view_count"`
	ClickCount         int    `json:"click_count"`
	CreateTime         int64  `json:"create_time"`
	CreateTimeStr      string `json:"create_time_str"`
	VisitorAlbumStatus int    `json:"visitor_album_status"`
	VisitorAlbumID     int    `json:"visitor_album_id"`
	MediaInfo          *Media `json:"media_info"`
}

type Media struct {
	MediaID      int             `json:"media_id"`
	UserID       int             `json:"user_id"`
	FileType     int             `json:"file_type"`
	FileURL      string          `json:"file_url"`
	FileName     string          `json:"file_name"`
	FileExt      string          `json:"file_ext"`
	FileSize     int             `json:"file_size"`
	Width        int             `json:"width"`
	Height       int             `json:"height"`
	DeletedImage string          `json:"deleted_image"`
	ExifInfo     json.RawMessage `json:"exif_info"`
}
