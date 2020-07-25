package models

const (
	DefaultSavePath = "data"
)

type CommonFlags struct {
	SavePath string
	Retry    int
}
