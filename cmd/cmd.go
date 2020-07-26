package cmd

import (
	"github.com/jdxj/wallpaper/models"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "wallpaper",
}

var commFlags = &models.CommonFlags{}

func init() {
	rootCmd.PersistentFlags().StringVar(&commFlags.SavePath, "path", "data", "set save path")
	rootCmd.PersistentFlags().IntVar(&commFlags.Concurrent, "concurrent", 10, "set goroutine pool size")
	rootCmd.PersistentFlags().IntVar(&commFlags.Retry, "retry", 3, "set retry times")
}

func Execute() error {
	return rootCmd.Execute()
}
