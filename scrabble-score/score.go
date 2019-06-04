package scrabble

import "unicode"

// Score calculates scrabble score
func Score(word string) int {
	var total int
	for _, v := range word {
		switch unicode.ToUpper(v) {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'R', 'N', 'S', 'T':
			total++
		case 'G', 'D':
			total += 2
		case 'B', 'C', 'M', 'P':
			total += 3
		case 'F', 'H', 'V', 'Y', 'W':
			total += 4
		case 'K':
			total += 5
		case 'J', 'X':
			total += 8
		case 'Q', 'Z':
			total += 10
		}

	}
	return total

}
