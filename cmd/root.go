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

	"github.com/gsilvapt/pmz/internal/logs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile      string
	ztldir       string
	notetemplate string
	editor       string
	gitrepo      string
	gituser      string
	repotoken    string
	Logger       *logs.Log
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pmz",
	Short: "Poor Man Zettelkasten CLI",
	Long: `This is a simple CLI application to help users maintain a Zettelkasten.
It provides methods to add, search and save your changes into a git repository.

Full documentation can be found in the project's README: https://github.com/gsilvapt/pmz
`,
	Version: "0.4",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	Logger = logs.InitLogger()
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pmz.yaml)")
	rootCmd.PersistentFlags().StringVar(&ztldir, "ztldir", "", "Zettelkasten main directory loaded from config.")
	rootCmd.PersistentFlags().StringVar(&notetemplate, "notetemplate", "", "Path to template of a new note. Ensure it contains the variables specified in the documentation.")
	rootCmd.PersistentFlags().StringVar(&editor, "editor", "", "Editor PATH loaded from config file.")
	rootCmd.PersistentFlags().StringVar(&gitrepo, "gitrepo", "", "Zettelkasten git repository loaded from config.")
	rootCmd.PersistentFlags().StringVar(&gituser, "gituser", "", "Git username loaded from config.")
	rootCmd.PersistentFlags().StringVar(&repotoken, "repotoken", "", "Zettelkasten git repository token with **read** and **write** access")

	rootCmd.InitDefaultVersionFlag()

	// Viper binding for global reach
	viper.BindPFlag("ztldir", rootCmd.PersistentFlags().Lookup("ztldir"))
	viper.BindPFlag("editor", rootCmd.PersistentFlags().Lookup("editor"))
	viper.BindPFlag("notetemplate", rootCmd.PersistentFlags().Lookup("notetemplate"))
	viper.BindPFlag("gitrepo", rootCmd.PersistentFlags().Lookup("gitrepo"))
	viper.BindPFlag("gituser", rootCmd.PersistentFlags().Lookup("gituser"))
	viper.BindPFlag("repotoken", rootCmd.PersistentFlags().Lookup("repotoken"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".pmz.yaml"
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".pmz.yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		Logger.Error(fmt.Sprintf("Failed reading config file: %s", viper.ConfigFileUsed()))
	}
}
