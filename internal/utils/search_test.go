package utils

import (
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
