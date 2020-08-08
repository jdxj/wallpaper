package user

import "testing"

func TestInfo_Query(t *testing.T) {
	flags := &Flags{
		ID:       "201438662",
		Identity: true,
	}
	info := NewInfo(flags)
	info.Query()
}
