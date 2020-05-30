package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "wallpaper",
}

func Execute() error {
	return rootCmd.Execute()
}
