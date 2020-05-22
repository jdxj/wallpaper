package cache

import (
	"fmt"
	"testing"
)

func TestNewSQLite(t *testing.T) {
	if err := sqlite.Ping(); err != nil {
		t.Fatalf("%s\n", err)
	}
	defer sqlite.Close()

	if err := SaveValue(Wallhaven, "ab", "cd"); err != nil {
		t.Fatalf("%s\n", err)
	}
}

func TestIsVisited(t *testing.T) {
	visited, err := IsVisited(Wallhaven, "w")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Println(visited)
}
