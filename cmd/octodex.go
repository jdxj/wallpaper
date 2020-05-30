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
	"github.com/jdxj/wallpaper/octodex"

	"github.com/spf13/cobra"
)

// octodexCmd represents the octodex command
var octodexCmd = &cobra.Command{
	Use:   "octodex",
	Short: "Download octodex avatar",
	Run: func(cmd *cobra.Command, args []string) {
		oc := octodex.NewCrawler(octFlags)
		oc.PushURL()
	},
}

var octFlags = &octodex.Flags{}

func init() {
	rootCmd.AddCommand(octodexCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// octodexCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// octodexCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	octodexCmd.Flags().StringVarP(&octFlags.Path, "path", "p", octodex.Path,
		"path specifies the storage path of the picture")
}
