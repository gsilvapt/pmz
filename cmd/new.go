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

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "A new note.",
	Long: `Creates a new note and opens it in your $EDITOR. 
For example:

<pmz new> will create a new note (README.md) in a directory with the current timestamp and opens it up in your $EDITOR.
You will see the new directory and file created in the configured ZTLDIR.

If a template is specified and you use the --title flag, it will try to insert the title in that template.
`,
	Run: func(cmd *cobra.Command, args []string) {
		ztldir := viper.GetString("ztldir")
		editor := viper.GetString("editor")
		tmpl_path := viper.GetString("notetemplate")
		title, _ := cmd.Flags().GetString("title")
		toOpen, _ := cmd.Flags().GetBool("open")

		ts := time.Now()
		formattedTs := fmt.Sprintf("%d%s%s%s%s%s", ts.Year(), ts.Format("01"), ts.Format("02"),
			ts.Format("03"), ts.Format("04"), ts.Format("05"),
		)

		f := createFile(ztldir, formattedTs, title)
		defer f.Close()

		if tmpl_path != "" {
			writeTmplToNote(f, title, tmpl_path)
		}

		if toOpen {
			OpenFile(f.Name(), editor)
		}
	},
}

// createFile attempts to create the directory for the new note, as well as a README.md file in that same directory.
// Panics if it any of it fails as this is a key functionality of the command.
// Returns file pointer. Caller should use `defer` to close this file pionter, or manually close it when not needed.
func createFile(ztldir, dirname, title string) *os.File {
	newZtlDir := fmt.Sprintf("%s/%s", ztldir, dirname)
	if err := os.Mkdir(newZtlDir, os.ModePerm); err != nil {
		Logger.Error(fmt.Sprintf("failed to create directory: %s", err))
	}

	var filename string = "README.md"
	if title != "" {
		filename = fmt.Sprintf("%s.md", title)
	}

	f, err := os.Create(fmt.Sprintf("%s/%s", newZtlDir, filename))
	if err != nil {
		Logger.Error(fmt.Sprintf("failed to create the file in new dir: %s", err))
	}

	return f
}

// writeTmplToNote writes the template into the new file. Handles errors gracefully.
func writeTmplToNote(f *os.File, title, tmplPath string) {
	tmplData := struct {
		Title string
	}{
		Title: title,
	}

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		Logger.Error("template file might not exist; not writing to new note.")
	}

	if err := tmpl.Execute(f, tmplData); err != nil {
		Logger.Info("Failed to write template to new note, but the note should exist.")
	}
	return
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().String("title", "", "Provide a title for the new note")
	newCmd.Flags().Bool("open", true, "To open or not the new note with the configured $EDITOR. Default is true.")
}
