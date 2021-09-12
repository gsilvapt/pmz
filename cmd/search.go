/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
    "os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// searchCmd represents the search command
// TODO Should allow searching by day, month and/or year (with the right flags)
// TODO Allow passing flags to grep directly for custom searches
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Searches for given keywords.",
	Long: `Searches for keywords in all Zettelkasten's notes and files. It integrates Grep and returns its output to 
    the main screen.`,
	Run: func(cmd *cobra.Command, args []string) {
		ztldir := viper.GetString("ztldir")
		term, _ := cmd.Flags().GetString("term")

        grep := exec.Command("grep", "-rH", "--exclude-dir", ".git", ztldir, "-e", term)
        grep.Stdin = os.Stdin
        grep.Stdout = os.Stdout
        grep.Stderr = os.Stderr
        grep.Run()
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")
	searchCmd.PersistentFlags().String("term", "", "The search term")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	searchCmd.Flags().IntP("day", "d", 0, "Look for messages created at day D - accepts from 1 to 31.")
	searchCmd.Flags().IntP("month", "m", 0, "Look for messages created at month M - accepts from 1 to 12.")
	searchCmd.Flags().IntP("year", "y", 0, "Look for messages created at year Y - accepts all positives numbers.")
}
