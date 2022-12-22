package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func ExecuteCmdWithArgsInTest(t *testing.T, cmd *cobra.Command, args ...string) (string, error) {
	t.Helper()

	buffer := new(bytes.Buffer)
	cmd.SetOut(buffer)
	cmd.SetErr(buffer)
	cmd.SetArgs(args)

	err := cmd.Execute()
	return strings.TrimSpace(buffer.String()), err
}
