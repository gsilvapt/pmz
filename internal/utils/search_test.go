package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestValidateSupportedExtensions(t *testing.T) {
	cases := map[string]bool{
		"avalid.md":         true,
		"invalid":           false,
		"alsosupported.rst": true,
		"alsosupported.txt": true,
	}

	for term, expected := range cases {
		actual := supportedExtension(filepath.Ext(term))
		if actual != expected {
			t.Fail()
		}
	}
}

func TestSearchInEmptyDirAndWithSearchableContent(t *testing.T) {
	// GIVEN an empty dir
	emptyTmpDir := t.TempDir()

	// OR GIVEN a temp dir with some notes
	tmpDir := t.TempDir()
	var searchableNote string = "Some searchable text\n"
	noteContents := []string{searchableNote, "Text it will not appear\n"}
	for idx, note := range noteContents {
		os.WriteFile(fmt.Sprintf("%s/note%d.md", tmpDir, idx), []byte(note), 0644)
	}

	// WHEN WalkNoteDir is called with the empty dir
	emptyResults := WalkNoteDir("search", emptyTmpDir)

	// THEN the result array is empty
	if len(emptyResults) != 0 {
		t.Fatalf("expected emptyResults array to be empty but it contains %d in it", len(emptyResults))
	}

	// BUT WHEN WalkNoteDir is called with the temp dir with some notes
	foundResults := WalkNoteDir("search", tmpDir)

	// THEN the result array is not empty and only contains the note with matching string
	if len(foundResults) != 1 {
		t.Fatalf("expected found results slice to have 1 entry but it contains %d", len(foundResults))
	}

	if foundResults[0].Context != searchableNote {
		t.Fatalf("expected \"%s\" to be note's context, but had %s", searchableNote, foundResults[0].Context)
	}
}
