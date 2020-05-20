package cmd

import (
	"fmt"
	"os"

	"github.com/jdxj/wallpaper/polayoutu"

	"github.com/jdxj/wallpaper/octodex"

	"github.com/jdxj/wallpaper/darenyou"
)

var (
	cmdHandlerSet map[string]Handler
)

func init() {
	cmdHandlerSet = make(map[string]Handler)

	RegisterCmdHandler(darenyou.NewCmdParser())
	RegisterCmdHandler(octodex.NewCmdParser())
	RegisterCmdHandler(polayoutu.NewCmdParser())
}

type Handler interface {
	ParseCmd([]string) error
	Run()
}

func Usage() {
	_, _ = fmt.Fprintf(os.Stderr, `Usage of wallpaper:
    wallpaper <command> [arguments]

Command:
    octodex
    polayoutu
    darenyou

Example:
    wallpaper polayoutu -h
`)
}

func RegisterCmdHandler(flagName string, h Handler) {
	if _, ok := cmdHandlerSet[flagName]; ok {
		err := fmt.Errorf("this flag name already exists: %s", flagName)
		panic(err)
	}

	if h == nil {
		err := fmt.Errorf("this handler is nil")
		panic(err)
	}

	cmdHandlerSet[flagName] = h
}

func HandleCmd(flagName string) {
	h, ok := cmdHandlerSet[flagName]
	if !ok {
		_, _ = fmt.Fprintf(os.Stderr, "wrong subcommand: %s\n",
			flagName)

		Usage()
		return
	}

	if err := h.ParseCmd(os.Args[2:]); err != nil {
		fmt.Printf("parse cmd err: %s\n", err)
		return
	}

	h.Run()
}
