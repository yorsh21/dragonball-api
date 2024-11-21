package utils

import (
	"strings"
)

func CapitalizeWords(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
	}
	return strings.Join(words, " ")
}
