package utils

import (
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var SupportedExtensions = []string{".md", ".txt", ".rst"}

type Result struct {
	Path    string
	Context string
}

// WalkNoteDir looks for supported files in the provided directory. Returns a list of Results if any found.
func WalkNoteDir(searchTerm string, path string) []*Result {
	var results []*Result
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.Contains(path, ".git") { // skip .git directory
			return nil
		}

		if !supportedExtension(filepath.Ext(path)) {
			return nil
		}

		f, err := os.OpenFile(path, os.O_RDONLY, 0600)
		if err != nil {
			return err
		}
		defer f.Close()

		rd := bufio.NewReader(f)
		for i := 0; i <= 1; i++ {
			line, err := rd.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					break
				}
				log.Fatal("failed reading line from file failed")
			}

			if strings.Contains(strings.ToLower(line), strings.ToLower(searchTerm)) {
				results = append(results, &Result{Path: path, Context: line})
			}
		}

		return nil
	})

	return results
}

func supportedExtension(term string) bool {
	for _, v := range SupportedExtensions {
		if v == term {
			return true
		}
	}

	return false
}
