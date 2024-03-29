package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	h := make(map[rune]int)
	h['A'] = 0
	h['C'] = 0
	h['G'] = 0
	h['T'] = 0
	for _, c := range d {
		if (c != 'A' && c != 'C' && c != 'G' && c != 'T') {
			return nil, errors.New("invalid character")
		}
		h[c]++
	}
	return Histogram(h), nil
}
