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
package poco

import (
	"github.com/jdxj/wallpaper/app/poco"
	"github.com/jdxj/wallpaper/cmd"

	"github.com/spf13/cobra"
)

// pocoCmd represents the poco command
var pocoCmd = &cobra.Command{
	Use:   "poco",
	Short: "poco",
	Run: func(cmd *cobra.Command, args []string) {
		poco.Run(pocoFlags)
	},
}

var pocoFlags = &poco.Flags{
	CommonFlags: cmd.CommFlags,
}

func init() {
	cmd.RootCmd.AddCommand(pocoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pocoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pocoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	pocoCmd.Flags().StringVarP(&pocoFlags.UserID, "userID", "u", "", "specify user id")
	pocoCmd.Flags().IntVarP(&pocoFlags.WorkID, "workID", "w", 0, "specify work id")
}
