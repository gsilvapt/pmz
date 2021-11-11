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

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Pushes changes to remote repository",
	Long: `Users can configure a remote git repository that allows them to "save" their notes onto the web.
    Whether that is BitBucket, GitHub, GitLab or anything else, it should work the same **as long as** the provided 
    token has write access to that repository and that repository already exists.`,
	Run: runSaveCommand,
}

func runSaveCommand(cmd *cobra.Command, args []string) {
	ztlrepo := viper.GetString("ztlrepo")
	gituser := viper.GetString("gituser")
	gittoken := viper.GetString("repotoken")

	r, err := git.PlainOpen(ztlrepo)
	if err != nil {
		Logger.Error(fmt.Sprintf("failed opening the git repository: %s", err))
	}

	wt, err := r.Worktree()
	if err != nil {
		Logger.Error(fmt.Sprintf("failed opening the repository tree: %s", err))
	}

	status, err := wt.Status()
	if err != nil {
		Logger.Error(fmt.Sprintf("failed getting the repository status: %s", err))
	}

	for file, status := range status {
		if status.Worktree == 'M' || status.Worktree == '?' {
			wt.Add(file)
		}
	}

	_, err = wt.Commit("updating your zettelkasten", &git.CommitOptions{})
	if err != nil {
		Logger.Error(fmt.Sprintf("failed committing changes: \n %s", err))
	}

	if err := r.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth: &http.BasicAuth{
			Username: gituser,
			Password: gittoken,
		},
	}); err != nil {
		Logger.Error(fmt.Sprintf("failed pushing to remote repository: \n %s", err))
	}

	Logger.Info("successfully saved your local changes to the remote repository.")
}

func init() {
	rootCmd.AddCommand(saveCmd)
}
