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
				if err == io.EOF {
					continue
				}
				log.Fatalf("failed reading first line from file: %s", err)
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
