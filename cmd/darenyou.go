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
	"github.com/jdxj/wallpaper/darenyou"

	"github.com/spf13/cobra"
)

// darenyouCmd represents the darenyou command
var darenyouCmd = &cobra.Command{
	Use:   "darenyou",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		darenyou.Run(dryFlags)
	},
}

var dryFlags = &darenyou.Flags{}

func init() {
	rootCmd.AddCommand(darenyouCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// darenyouCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// darenyouCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	darenyouCmd.Flags().StringVarP(&dryFlags.SavePath, "savePath", "a", "data", "set save path")
	darenyouCmd.Flags().IntVarP(&dryFlags.Retry, "retry", "r", 3, "set retry times")

	darenyouCmd.Flags().StringVarP(&dryFlags.Project, "project", "p", darenyou.Chaos, "select a project (photo album)")
	darenyouCmd.Flags().StringVarP(&dryFlags.Size, "size", "s", darenyou.Src, "specify picture resolution")
}
