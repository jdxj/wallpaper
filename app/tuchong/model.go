package tuchong

import (
	"compress/gzip"
	"encoding/json"
	"io"
)

type FollowResponse struct {
	Message        string `json:"message"`
	Hint           string `json:"hint"`
	ShowPointToast bool   `json:"showPointToast"`
	Point          int    `json:"point"`
	Result         string `json:"result"`
}

func UnmarshalFollowResponse(body io.ReadCloser, isGzip bool) (*FollowResponse, error) {
	defer body.Close()

	reader := body.(io.Reader)
	if isGzip {
		r, err := gzip.NewReader(reader)
		if err != nil {
			return nil, err
		}
		reader = r
	}

	decoder := json.NewDecoder(reader)
	resp := &FollowResponse{}
	return resp, decoder.Decode(resp)
}

type Works struct {
	Count           int     `json:"count"`
	WorkList        []*Work `json:"work_list"`
	BeforeTimestamp int64   `json:"before_timestamp"`
	More            bool    `json:"more"`
	Result          string  `json:"result"`
}

type Work struct {
}
