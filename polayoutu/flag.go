package polayoutu

type Flags struct {
	Path    string
	Size    string
	Edition int
}

const (
	Path = "data"

	FullRes = "full"
	Thumb   = "thumb"
	Size    = Thumb

	EditionNum = 1
)
