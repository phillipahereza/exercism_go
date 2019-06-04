package strand

var complements = map[rune]rune{
	'A': 'U',
	'G': 'C',
	'C': 'G',
	'T': 'A',
}

// ToRNA transcribes a DNA string to RNA
func ToRNA(dna string) (rna string) {
	for _, r := range dna {
		rna += string(complements[r])
	}
	return rna
}
