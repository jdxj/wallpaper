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
	"github.com/jdxj/wallpaper/app/androidesk/album"

	"github.com/spf13/cobra"
)

// albumCmd represents the album command
var albumCmd = &cobra.Command{
	Use:   "album",
	Short: "download androidesk album",
	Run: func(cmd *cobra.Command, args []string) {
		album.Run(albumFlags)
	},
}

var albumFlags = &album.Flags{
	CommonFlags: commFlags,
}

func init() {
	androideskCmd.AddCommand(albumCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// albumCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// albumCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	albumCmd.Flags().StringVarP(&albumFlags.ID, "id", "n", "5d834a6fe7bce73981fabf4c", "album id. the default value is only for testing")
	albumCmd.Flags().IntVarP(&albumFlags.Limit, "limit", "l", 0, "specify the number of download pages")
	albumCmd.Flags().BoolVarP(&albumFlags.Adult, "adult", "a", false, "may not have adult content")
	albumCmd.Flags().StringVarP(&albumFlags.Order, "order", "o", album.New, "specify collation")
}
