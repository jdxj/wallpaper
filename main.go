package main

import (
	"fmt"

	"github.com/jdxj/wallpaper/cmd"
	"github.com/jdxj/wallpaper/db"
	"github.com/jdxj/wallpaper/utils"
)

func main() {
	go func() {
		utils.ReceiveInterrupt()
	}()

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}

	db.Close()
}
