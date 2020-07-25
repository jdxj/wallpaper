package models

const (
	DefaultSavePath = "data"
)

type Flags struct {
	SavePath string
	Retry    int
}
