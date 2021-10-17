/*
Copyright © 2021 GUSTAVO SILVA <gustavosantaremsilva@gmail.com>

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
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const FILENAME string = "README.md"

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "A new note.",
	Long: `Creates a new note and opens it in your $EDITOR. 
For example:

<pmz new> will create a new note (README.md) in a directory with the current timestamp and opens it up in your $EDITOR.
You will see the new directory and file created in the configured ZTLDIR.
`,
	Run: func(cmd *cobra.Command, args []string) {
		ztldir := viper.GetString("ztldir")
		editor := viper.GetString("editor")
		title, _ := cmd.Flags().GetString("title")
		toOpen, _ := cmd.Flags().GetBool("open")

		ts := time.Now()
		formattedTs := fmt.Sprintf("%d%d%d%d%d", ts.Year(), ts.Month(), ts.Day(), ts.Hour(), ts.Minute())

		f := createZettelEntry(ztldir, formattedTs)
		defer f.Close()

		writeToNewNote(f, title)
		if toOpen {
			OpenFile(f.Name(), editor)
		}
	},
}

// createZettelEntry attempts to create the directory for the new note, as well as a README.md file in that same directory.
// Panics if it any of it fails as this is a key functionality of the command.
// Returns file pointer. Caller should use `defer` to close this file pionter, or manually close it when not needed.
func createZettelEntry(ztldir, dirname string) *os.File {
	newZtlDir := fmt.Sprintf("%s/%s", ztldir, dirname)

	if err := os.Mkdir(newZtlDir, os.ModePerm); err != nil {
		PanicIfError(err, "Failed to create directory.")
	}

	f, err := os.Create(fmt.Sprintf("%s/%s", newZtlDir, FILENAME))
	if err != nil {
		PanicIfError(err, "Failed to create the file in new dir.")
	}

	return f
}

// writeToNewNote writes the template into the new file. Handles errors gracefully.
func writeToNewNote(f *os.File, title string) {
	tmplData := struct {
		Title string
	}{
		Title: title,
	}

	tmpl, err := template.ParseFiles("templates/new_note")
	if err != nil {
		fmt.Println("Failed to write template to new note, but it should exist for you.")
	}

	if err := tmpl.Execute(f, tmplData); err != nil {
		fmt.Println("Failed to write template to new note, but it should exist for you.")
	}
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().String("title", "", "Provide a title for the new note")
	newCmd.Flags().Bool("open", true, "To open or not the new note with the configured $EDITOR. Default is true.")
}
