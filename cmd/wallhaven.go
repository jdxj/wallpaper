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
		wc := wallhaven.NewCrawler(walFlags)
		wc.PushURL()
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

	wallhavenCmd.Flags().BoolVar(&walFlags.General, "general", false,
		"set general categories bit")
	wallhavenCmd.Flags().BoolVar(&walFlags.Anime, "anime", false,
		"set anime categories bit")
	wallhavenCmd.Flags().BoolVar(&walFlags.People, "people", false,
		"set people categories bit")

	wallhavenCmd.Flags().BoolVar(&walFlags.Sfw, "sfw", false,
		"set sfw purity bit")
	wallhavenCmd.Flags().BoolVar(&walFlags.Sketchy, "sketchy", false,
		"set sketchy purity bit")
	wallhavenCmd.Flags().BoolVar(&walFlags.Nsfw, "nsfw", false,
		"set nsfw purity bit")

	wallhavenCmd.Flags().StringVar(&walFlags.Sorting, "sorting", wallhaven.TopList,
		"set collation "+wallhaven.SortingOptionalValue)
	wallhavenCmd.Flags().StringVar(&walFlags.TopRange, "topRange", "",
		"set time range "+wallhaven.TopRangeOptionalValue)
	wallhavenCmd.Flags().StringVar(&walFlags.Order, "order", wallhaven.Desc,
		"set order "+wallhaven.OrderOptionalValue)
	wallhavenCmd.Flags().IntVar(&walFlags.Page, "page", 1,
		"set page number, does not limit the scope")

	wallhavenCmd.Flags().StringVarP(&walFlags.Path, "path", "p", wallhaven.Path,
		"set storage path")
	wallhavenCmd.Flags().StringVarP(&walFlags.Url, "url", "u", "",
		"specify the url of the image to download")
}
