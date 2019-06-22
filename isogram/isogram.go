// Package isogram provides a utility for determining
// if a word or phrase is an isogram.
package isogram

import (
	"unicode"
)

// IsIsogram returns true if string is isogram.
func IsIsogram(s string) bool {
	seen := map[rune]bool{}
	for _, char := range s {
		char := unicode.ToLower(char)
		if unicode.IsLetter(char) {
			if _, ok := seen[char]; ok {
				return false
			}
			seen[char] = true
		}
	}

	return true
	
}
