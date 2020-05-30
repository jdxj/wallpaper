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
	"github.com/jdxj/wallpaper/polayoutu"

	"github.com/spf13/cobra"
)

// polayoutuCmd represents the polayoutu command
var polayoutuCmd = &cobra.Command{
	Use:   "polayoutu",
	Short: "Download photos from polayoutu",
	Run: func(cmd *cobra.Command, args []string) {
		pc := polayoutu.NewCrawler(polFlags)
		pc.PushURL()
	},
}

var polFlags = &polayoutu.Flags{}

func init() {
	rootCmd.AddCommand(polayoutuCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// polayoutuCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// polayoutuCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	polayoutuCmd.Flags().StringVarP(&polFlags.Path, "path", "p", polayoutu.Path,
		"path specifies the storage path of the picture")
	polayoutuCmd.Flags().StringVarP(&polFlags.Size, "size", "s", polayoutu.Size,
		"size specifies the resolution of the image to be downloaded. you can choose [full | thumb]")
	polayoutuCmd.Flags().IntVarP(&polFlags.Edition, "edition", "e", polayoutu.EditionNum,
		"edition specifies which period to download")
}