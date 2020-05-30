package main

import (
	"fmt"

	"github.com/jdxj/wallpaper/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
