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

var upToDateErrorMessage string = "already up-to-date"

// updateCmd represents the update command
// TODO pull changes from git repository (if it is configured in the file).
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Pulls data from remote repository",
	Long: `Users can configure a remote git repository that allows them to pull changes from there to the local 
    copy. This feature allows users to use multiple devices to build their own Zettelkasten.

    This command should work **as long as** the provided token has read access to that repository and that 
    repository already exists remotely.`,
	Run: runUpdateCmd,
}

func runUpdateCmd(cmd *cobra.Command, args []string) {
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

	switch err := wt.Pull(&git.PullOptions{
		RemoteName: "origin",
		Auth: &http.BasicAuth{
			Username: gituser,
			Password: gittoken,
		},
	}); err != nil {
	case err.Error() == upToDateErrorMessage:
		Logger.Info("repository already up to date")
		return
	default:
		Logger.Error(fmt.Sprintf("failed pulling from remote repository: \n %s", err))
	}

	Logger.Info("Successfully updated your local Zettelkasten.")
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
