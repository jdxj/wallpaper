/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/jdxj/wallpaper/wallhaven"

	"github.com/spf13/cobra"
)

// wallhavenCmd represents the wallhaven command
var wallhavenCmd = &cobra.Command{
	Use:   "wallhaven",
	Short: "Download wallpapers from wallhaven",
	Run: func(cmd *cobra.Command, args []string) {
		wallhaven.Run(walFlags)
	},
}

var walFlags = &wallhaven.Flags{}

func init() {
	rootCmd.AddCommand(wallhavenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wallhavenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wallhavenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	wallhavenCmd.Flags().StringVarP(&walFlags.SavePath, "savePath", "sp", "data", "set save path")
	wallhavenCmd.Flags().IntVarP(&walFlags.Retry, "retry", "r", 3, "set retry times")

	wallhavenCmd.Flags().StringVarP(&walFlags.UserName, "userName", "un", "", "set user name")
	wallhavenCmd.Flags().StringVarP(&walFlags.CollectionID, "collectionID", "ci", "", "set collection id")
	wallhavenCmd.Flags().StringVarP(&walFlags.APIKey, "apiKey", "ak", "", "set api key")
}
