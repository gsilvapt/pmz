package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func PanicIfError(err error, msg string) {
	if err != nil {
		panic(fmt.Sprintf("%s: \n %s", msg, err))
	}
}

// OpenFile opens the specified file in the editor that both are provided.
// Will likely fail if file does not exist, or path to editor is poor. Callee should ensure both are correct.
func OpenFile(path, editor string) {
	cmd := exec.Command(editor, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
