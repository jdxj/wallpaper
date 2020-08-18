package cmd

import (
	"github.com/jdxj/wallpaper/models"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "wallpaper",
}

var CommFlags = &models.CommonFlags{}

func init() {
	RootCmd.PersistentFlags().StringVar(&CommFlags.SavePath, "path", "data", "set save path")
	RootCmd.PersistentFlags().IntVar(&CommFlags.Concurrent, "concurrent", 10, "set goroutine pool size")
	RootCmd.PersistentFlags().IntVar(&CommFlags.Retry, "retry", 3, "set retry times")
	RootCmd.PersistentFlags().IntVar(&CommFlags.Timeout, "timeout", 30, "client read http body timeout")
}

func Execute() error {
	return RootCmd.Execute()
}
