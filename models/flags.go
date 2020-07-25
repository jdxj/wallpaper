package models

const (
	DefaultSavePath = "data"
)

type Config struct {
	SavePath string
	Retry    int
}
