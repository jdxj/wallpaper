package utils

import (
	"os"
	"testing"
)

func TestWriteToFile(t *testing.T) {
	if err := WriteToFile("abc/dea", []byte("abc")); err != nil {
		t.Fatalf("%s", err)
	}
}

func TestWriteToFile2(t *testing.T) {
	if err := os.MkdirAll("abc/def.t", os.ModePerm); err != nil {
		t.Fatalf("%s", err)
	}
}
