package cmd

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func TestSearchCMDIsAbleToSearchNotesTitle(t *testing.T) {
	// GIVEN a ztldir with two notes
	searchCmd := &cobra.Command{Use: "search command", Run: searchCmdFunc}

	// set ztldir to controlled directory for better matching
	tmpDir := t.TempDir()

	viper.Set("ztldir", tmpDir)
	// Write fake content to notes
	var searchableNote string = "Some searchable text\n"
	noteContents := []string{searchableNote, "Text that will not appear\n"}
	for idx, note := range noteContents {
		os.WriteFile(fmt.Sprintf("%s/note%d.md", tmpDir, idx), []byte(note), 0644)
	}

	// WHEN user searches for any word
	output, err := ExecuteCmdWithStdinPipe(t, searchCmd, "q\n", "search")

	// THEN the pmz search returns a single note with the searchable content.
	if err != nil {
		t.Fatalf("command unexpectably errored out: %s", err)
	}

	expectedOut := strings.TrimRight(fmt.Sprintf("0 | %s/note0.md: %s", tmpDir, searchableNote), "\n")
	outLines := strings.Split(output, "\n")
	actualOut := outLines[0]

	if expectedOut != actualOut {
		t.Fatalf("expected output to contain note title:\n%s\nand instead got the following:\n%s", expectedOut, actualOut)
	}
}

func TestSearchCMDInEmptyDirReturnsNoResults(t *testing.T) {
	// GIVEN an empty ztldir with no notes
	searchCmd := &cobra.Command{Use: "search command", Run: searchCmdFunc}

	// set ztldir to controlled directory for better matching
	tmpDir := t.TempDir()
	viper.Set("ztldir", tmpDir)

	// WHEN user searches for any word
	output, err := ExecuteCmdWithStdoutPipe(t, searchCmd, "search")

	// THEN the pmz search returns a single note with the searchable content.
	if err != nil {
		t.Fatalf("command unexpectably errored out: %s", err)
	}

	outLines := strings.Split(output, "\n")
	actualOut := outLines[0]

	expectedOut := "No results found for query. Exiting..."
	if expectedOut != actualOut {
		t.Fatalf("expected output to be:\n%s\nbut instead got:\n%s", expectedOut, actualOut)
	}
}
