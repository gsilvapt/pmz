package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func TestNewCmdWithNotOpen(t *testing.T) {
	// GIVEN a _new_ Cobra CMD sruct with the function newCmdRunFunc in the struct
	newCommand := &cobra.Command{Use: "new command", Run: newNoteFunc}

	// AND GIVEN a viper config with disabled toOpen and a random ztldir attribute
	tmpDir := t.TempDir()
	viper.Set("ztldir", tmpDir)

	// AND GIVEN the ztldir is empty
	files, err := ioutil.ReadDir(tmpDir)
	if err != nil {
		t.Fatalf("failed to read from temporary directory for tests: %s", err)
	}

	if len(files) != 0 {
		t.Fatalf("temporary directory contains %d files when it should be zero!!", len(files))
	}

	// WHEN the command is called and runs the newCmdRunFunc with the desired args
	PrepNewCmdFlags(newCommand)
	_, err = ExecuteCmdWithArgsInTest(t, newCommand, "--open=false")
	if err != nil {
		t.Fatalf("command unexpectably errored out: %s", err)
	}

	// THEN a new directory in it is created, containing a README.md file.
	files, err = ioutil.ReadDir(tmpDir)
	if err != nil {
		t.Fatalf("failed read from temporary directory: %s", err)
	}

	ztlDir := files[0].Name()
	files, err = ioutil.ReadDir(fmt.Sprintf("%s/%s", tmpDir, ztlDir))
	if err != nil {
		t.Fatalf("failed to read freshly new zettel dir: %s", err)
	}

	if files[0].Name() != "README.md" {
		t.Errorf("temporary directory should contain a single directory with a file in it with name README.md, but got: %d - %s", len(files), files[0].Name())
	}
}

func TestNewCmdWithTitle(t *testing.T) {
	// GIVEN a _new_ Cobra CMD sruct with the function newCmdRunFunc in the struct
	newCommand := &cobra.Command{Use: "new command", Run: newNoteFunc}

	// AND GIVEN a viper config with disabled toOpen and a random ztldir attribute
	tmpDir := t.TempDir()
	viper.Set("ztldir", tmpDir)

	// AND GIVEN the ztldir is empty
	files, err := ioutil.ReadDir(tmpDir)
	if err != nil {
		t.Fatalf("failed to read from temporary directory for tests: %s", err)
	}

	if len(files) != 0 {
		t.Fatalf("temporary directory contains %d files when it should be zero!!", len(files))
	}

	// WHEN the command is called and runs the newCmdRunFunc with the desired args
	var testTitle string = "test"
	PrepNewCmdFlags(newCommand)
	_, err = ExecuteCmdWithArgsInTest(t, newCommand, "--open=false", fmt.Sprintf("--title=%s", testTitle))
	if err != nil {
		t.Fatalf("command unexpectably errored out: %s", err)
	}

	// THEN a new directory in it is created, containing a file with the provided title.
	files, err = ioutil.ReadDir(tmpDir)
	if err != nil {
		t.Fatalf("failed read from temporary directory: %s", err)
	}

	var expectedFile string = fmt.Sprintf("%s/%s/%s.md", tmpDir, files[0].Name(), testTitle)
	file, err := os.Open(expectedFile)
	defer file.Close()
	if err != nil {
		t.Fatalf("failed reading expected output file %s: %s", expectedFile, err)
	}

	// AND THEN the first line of the file *contains* the provided title
	scanner := bufio.NewScanner(file)
	var line int = 0
	for scanner.Scan() {
		if line == 1 {
			actual := scanner.Text()
			t.Log(actual)
			if actual != testTitle {
				t.Errorf("expected first line of file to be %s but got %s", testTitle, actual)
				break
			}
		}
	}
}
