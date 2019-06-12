package isogram

import (
	"regexp"
	"strings"
)

// IsIsogram return true if string is isogram
func IsIsogram(s string) bool {
	
	freq := make(map[rune]rune)
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")

	processedString := reg.ReplaceAllString(strings.ToLower(s), "")
	for _, r := range processedString {
		freq[r] = r
	}

	return len(processedString) == len(freq)

}
