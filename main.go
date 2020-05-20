package main

import (
	"os"

	"github.com/jdxj/wallpaper/cmd"
)

func main() {
	if len(os.Args) < 2 {
		cmd.Usage()
		return
	}

	cmd.HandleCmd(os.Args[1])
}
