package search

const (
	Home = "home" // 综合
	User = "site" // 用户
)

type Flags struct {
	Type  string
	Query string
}
