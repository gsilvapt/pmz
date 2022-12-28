package cmd

import (
	"os"
	"os/exec"
)

// openFileInEditor opens the specified file in the editor that both are provided.
// Will likely fail if file does not exist, or path to editor is poor. Callee should ensure both are correct.
func openFileInEditor(path, editor string) {
	cmd := exec.Command(editor, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
