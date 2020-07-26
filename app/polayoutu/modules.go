package polayoutu

import "encoding/json"

type ResponseJson struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
	Code    string          `json:"code"`
}

type Edition struct {
	Edition string `json:"edition"`
}

type Photo struct {
	ID                    int    `json:"id"`
	UserID                int    `json:"user_id"`
	UserEmail             string `json:"user_email"`
	CollectionID          int    `json:"collection_id"`
	FullRes               string `json:"full_res"`
	Thumb                 string `json:"thumb"`
	Avatar                string `json:"avatar"`
	Story                 string `json:"story"`
	Upvotes               string `json:"upvotes"`
	UpvoteCount           int    `json:"upvote_count"`
	ImgColor              string `json:"img_color"`
	Created               string `json:"created"`
	CommentCount          int    `json:"comment_count"`
	WallpaperDisplayCount int    `json:"wallpaper_display_count"`
	Tags                  string `json:"tags"`
	UserName              string `json:"user_name"`
}
