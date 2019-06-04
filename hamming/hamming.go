package hamming

import "errors"

// Distance calculates the Hamming Difference between 2 DNA strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("strands must be of the same length")
	}
	hammingCount := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hammingCount++
		}
	}
	return hammingCount, nil
}
