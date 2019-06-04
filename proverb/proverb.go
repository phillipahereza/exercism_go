// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

const (
	verse = "For want of a %s the %s was lost."
	last  = "And all for the want of a %s."
)

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}
	poem := make([]string, len(rhyme))
	poem[len(rhyme)-1] = fmt.Sprintf(last, rhyme[0])

	prev := rhyme[0]

	for i, v := range rhyme[1:] {
		poem[i] = fmt.Sprintf(verse, prev, v)
		prev = v
	}

	return poem
}
