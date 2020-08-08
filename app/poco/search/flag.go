package search

const (
	User  = "user"
	Works = "works"
)

type Flags struct {
	Type    string
	Keyword string
	Page    int
}
