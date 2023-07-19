package exercism

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	filteredWord := strings.ReplaceAll(word, "-", "")
	filteredWord = strings.ReplaceAll(filteredWord, " ", "")

	uniqueChars := make(map[rune]bool)

	for _, char := range filteredWord {
		_, exists := uniqueChars[unicode.ToLower(char)]
		if exists {
			return false
		} else {
			uniqueChars[unicode.ToLower(char)] = true
		}
	}
	return true
}
