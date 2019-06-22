// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
	"regexp"
)



// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")

	processedString := reg.ReplaceAllString(s, " ")

	words := strings.Split(processedString, " ")
	abbr := make([]string, len(words))

	for _, word := range words {
		abbr = append(abbr, string([]rune(word)[0]))
	}

	return strings.ToUpper(strings.Join(abbr, ""))
}
