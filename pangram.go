package exercism

import "strings"

func IsPangram(input string) bool {
	chars := make(map[rune]bool)

	split := []rune(strings.ToLower(input))
	count := 0
	for _, char := range split {
		if char >= 'a' && char <= 'z' {
			if exists, _ := chars[char]; !exists {
				chars[char] = true
				count++
				if count == 26 {
					return true
				}
			}
		}
	}

	return false
}
