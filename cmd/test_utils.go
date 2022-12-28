package cmd

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

// ExecuteCmdWithArgsInTest is a wrapper to facilitate calling cobra.Command structs, in which arguments are passed
// to the Command struct and any potential output from those commands is captured to then be returned.
// This function **does not wrap os.Stdout nor os.Stderr**.
// To perform assertions over os.Stdout use ExecuteCmdWithPipes.
func ExecuteCmdWithArgsInTest(t *testing.T, cmd *cobra.Command, args ...string) (string, error) {
	t.Helper()

	buffer := new(bytes.Buffer)
	cmd.SetOut(buffer)
	cmd.SetErr(buffer)
	cmd.SetArgs(args)

	err := cmd.Execute()
	return strings.TrimSpace(buffer.String()), err
}

// ExecuteCmdWithStdoutPipe is a wrapper to facilitate calling cobra.Command structs, in which arguments are passed to the
// Command struct, but in this function os.Stdout is captured and returned to the caller.
// This function **wraps os.Stdout** but not os.Stderr.
// To perform assertions over Command returned values, use ExecuteCmdWithArgsInTest.
func ExecuteCmdWithStdoutPipe(t *testing.T, cmd *cobra.Command, args ...string) (string, error) {
	t.Helper()

	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	cmd.SetArgs(args)
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}

	w.Close()
	buffer, err := io.ReadAll(r)
	r.Close()

	os.Stdout = oldOut
	return strings.TrimSpace(string(buffer)), err
}

func ExecuteCmdWithStdinPipe(t *testing.T, cmd *cobra.Command, toStdin string, args ...string) (string, error) {
	t.Helper()

	oldOut := os.Stdout
	oldIn := os.Stdin

	inReader, inWriter, _ := os.Pipe()
	outReader, outWriter, _ := os.Pipe()

	os.Stdin = inReader
	os.Stdout = outWriter

	cmd.SetArgs(args)

	inWriter.Write([]byte(toStdin))
	cmd.Execute()

	inWriter.Close()
	inReader.Close()
	outWriter.Close()
	buffer, err := io.ReadAll(outReader)
	outReader.Close()

	os.Stdin = oldIn
	os.Stdout = oldOut

	return strings.TrimSpace(string(buffer)), err
}
