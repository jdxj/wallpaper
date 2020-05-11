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

var usableCmd = `usable cmd:
	octodex
	polayoutu
`

func main() {
	if len(os.Args) < 2 {
		fmt.Print(usableCmd)
		return
	}

	cmd := os.Args[1]
	f, ok := subCmd[cmd]
	if !ok {
		fmt.Printf("not found subcmd: %s\n", cmd)
		return
	}
	f()
}
