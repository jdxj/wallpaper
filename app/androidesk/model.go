package androidesk

import (
	"encoding/json"
	"io"
)

type Response struct {
	Message  string          `json:"msg"`
	Resource json.RawMessage `json:"res"`
	Code     int             `json:"code"`
}

func (resp *Response) Unmarshal(body io.ReadCloser) error {
	defer body.Close()

	decoder := json.NewDecoder(body)
	return decoder.Decode(resp)
}

type User struct {
	GcID      string  `json:"gcid"`
	Name      string  `json:"name"`
	Gender    int     `json:"gender"`
	Follower  int     `json:"follower"`
	Avatar    string  `json:"avatar"`
	VipTime   float64 `json:"viptime"`
	Following int     `json:"following"`
	IsVip     bool    `json:"isvip"`
	ID        string  `json:"id"`
	Auth      string  `json:"auth"`
}

// Wallpaper 元信息
type Wallpaper struct {
	Views   int      `json:"views"`
	NCos    int      `json:"ncos"`
	Rank    int      `json:"rank"`
	Tags    []string `json:"tag"`
	User    *User    `json:"user"`
	Wp      string   `json:"wp"`
	Xr      bool     `json:"xr"`
	Cr      bool     `json:"cr"`
	FAvs    int      `json:"favs"`
	Atime   float64  `json:"atime"`
	ID      string   `json:"id"`
	Desc    string   `json:"desc"`
	Thumb   string   `json:"thumb"`
	Img     string   `json:"img"`
	Cid     []string `json:"cid"`
	URL     []string `json:"url"`
	Rule    string   `json:"rule"`
	RuleNew string   `json:"rule_new"`
	Preview string   `json:"preview"`
	Store   string   `json:"store"`
}
