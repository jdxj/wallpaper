package darenyou

import "encoding/json"

type Project struct {
	ID         int             `json:"id"`
	SiteID     string          `json:"site_id"`
	ProjectURL string          `json:"project_url"`
	DirectLink string          `json:"direct_link"`
	Type       string          `json:"type"`
	Content    json.RawMessage `json:"content"`
}
