// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import ("strings"
"bytes"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	s = strings.ToUpper(s)
	var words = []string
	words = strings.Split(s, " ")
	var abbr bytes.Buffer
	for _, word := range(words) {
		word[0]
	}

	return abbr.String()
}
