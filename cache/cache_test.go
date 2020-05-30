package cache

import (
	"fmt"
	"testing"
)

func TestIsVisited(t *testing.T) {
	value, err := IsVisited([]byte("hello"))
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
	fmt.Printf("%s\n", value)
}

func TestSaveValue(t *testing.T) {
	key := []byte("hello")
	value := []byte("world")
	if err := SaveValue(key, value); err != nil {
		t.Fatalf("err: %s\n", err)
	}
}
