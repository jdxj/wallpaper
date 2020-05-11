package main

import (
	"fmt"
	"os"

	"github.com/jdxj/wallpaper/octodex"
	"github.com/jdxj/wallpaper/polayoutu"
)

var subCmd = map[string]func(){
	"octodex":   octodex.Run,
	"polayoutu": polayoutu.Run,
}

func Usage() {
	fmt.Fprintf(os.Stderr, `Usage of wallpaper:
    wallpaper <command> [arguments]

Command:
    octodex
    polayoutu
`)
}

func main() {
	if len(os.Args) < 2 {
		Usage()
		return
	}

	cmd := os.Args[1]
	f, ok := subCmd[cmd]
	if !ok {
		Usage()
		return
	}
	f()
}
