// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"
const minLetter, maxLetter = 'a', 'z'

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.Trim(remark, " \t\n\r")
	response := "Whatever."
	if strings.HasSuffix(remark, "?") && hasUppercase(remark) && hasLetters(remark) {
		response = "Calm down, I know what I'm doing!"
	} else if strings.HasSuffix(remark, "?") {
		response = "Sure."
	} else if remark == "" {
		response = "Fine. Be that way!"
	} else if hasUppercase(remark) && hasLetters(remark) {
		response = "Whoa, chill out!"
	}
	return response
}

func hasUppercase(s string) bool {
	return strings.ToUpper(s) == s
}

func hasLetters(s string) bool {
	toLower := strings.ToLower(s)
	for _, l := range toLower {
		if l <= maxLetter && l >= minLetter {
			return true
		}
	}
	return false
}
