package utils

import (
	"net/http"
	"os"
	"testing"
)

func TestWriteToFile2(t *testing.T) {
	if err := os.MkdirAll("abc/def.t", os.ModePerm); err != nil {
		t.Fatalf("%s", err)
	}
}
func TestDetectFileType(t *testing.T) {
	file, err := os.Open("image")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	defer file.Close()

	buf := make([]byte, 512)
	n, err := file.Read(buf)
	if n != 512 || err != nil {
		t.Fatalf("n: %d, err: %s\n", n, err)
	}

	result := http.DetectContentType(buf)
	t.Logf("file type: %s\n", result)
}

func TestWriteFromReadCloser(t *testing.T) {
	file, err := os.Open("image")
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	err = WriteFromReadCloser("data", "abc", file)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}
