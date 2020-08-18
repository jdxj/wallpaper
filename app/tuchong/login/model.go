package login

import (
	"compress/gzip"
	"encoding/json"
	"io"
)

type LoginResponse struct {
	Code    int    `json:"code"`
	Result  string `json:"result"`
	Message string `json:"message"`

	Identity string `json:"identity"`
	Token    string `json:"token"`
	Name     string `json:"name"`
	Icon     string `json:"icon"`

	Hint      string `json:"hint"`
	TagMarks  bool   `json:"tagmarks"`
	HasMobile bool   `json:"hasMobile"`
}

func UnmarshalLoginResponse(body io.ReadCloser, isGzip bool) (*LoginResponse, error) {
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
	resp := &LoginResponse{}
	return resp, decoder.Decode(resp)
}
