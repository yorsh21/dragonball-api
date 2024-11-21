package utils_test

import (
	"dragonball-api/pkg/utils"
	"testing"
)

func TestCapitalizeWords(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{"hello world", "Hello World"},
		{"hELLO wORLD", "Hello World"},
		{"single word", "Single Word"},
		{"", ""},
	}

	for _, tc := range testCases {
		result := utils.CapitalizeWords(tc.input)
		if result != tc.output {
			t.Errorf("CapitalizeWords(%q) = %q, want %q", tc.input, result, tc.output)
		}
	}
}
