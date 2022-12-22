package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gsilvapt/pmz/internal/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// searchCmd represents the search command
// TODO Should allow searching by day, month and/or year (with the right flags)
// TODO Allow passing flags to grep directly for custom searches
var searchCmd = &cobra.Command{
	Use:   "search <term>",
	Short: "Searches for given keywords.",
	Long: `Searches for keywords in all Zettelkasten's notes and files. It integrates Grep and returns its output to 
    the main screen.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ztldir := viper.GetString("ztldir")
		editor := viper.GetString("editor")
		term := args[0]

		// TODO Must make use of the outcome from WalkNoteDir and thus this recursive call should do something.
		var r []*utils.Result = utils.WalkNoteDir(term, ztldir)
		if len(r) < 1 {
			fmt.Println("No results found for query. Exiting...")
			return
		}

		for i, f := range r {
			fmt.Printf("%d | %s: %s", i, f.Path, f.Context)
		}

		// Proceed with next command
	Interaction:
		for {
			fmt.Println("Choose the next action with the found files: `open <id>` to open with your editor, " +
				"`more <id>` to print the file contents, or <q> to quit.")

			switch cmd, idx := nextCommand(); cmd {
			case "open":
				f := r[idx]
				OpenFile(f.Path, editor)
			case "more":
				f := r[idx]
				readFile(f.Path)
			case "q":
				break Interaction
			default:
				fmt.Println("Unrecognized option: `open <id>` to open with your editor, `more <id>`" +
					" to print file's contents or <q>  to quit.")
			}
		}
	},
}

func nextCommand() (string, int) {
	var command string = ""
	var idx int = 0

	buffer := bufio.NewReader(os.Stdin)
	line, err := buffer.ReadString('\n')
	if err != nil {
		Logger.Error(fmt.Sprintf("failed reading input from screen: %s", err))
	}

	switch inputs := strings.Fields(line); len(inputs) {
	case 1:
		command = inputs[0]
	case 2:
		idx, err = strconv.Atoi(inputs[1])
		if err != nil {
			Logger.Error(fmt.Sprintf("failed to detect number - did you provide one? %s", err))
		}

		command = inputs[0]
	case 0:
		command = ""
		idx = 0
	}

	return command, idx
}

func readFile(fp string) {
	dat, err := os.ReadFile(fp)
	if err != nil {
		Logger.Error(fmt.Sprintf("failed opening specified file: %s", err))
	}
	fmt.Println(string(dat))
}

func init() {
	searchCmd.Flags().IntP("day", "d", 0, "Look for messages created at day D - accepts from 1 to 31.")
	searchCmd.Flags().IntP("month", "m", 0, "Look for messages created at month M - accepts from 1 to 12.")
	searchCmd.Flags().IntP("year", "y", 0, "Look for messages created at year Y - accepts all positives numbers.")

	rootCmd.AddCommand(searchCmd)
}
