package accumulate

// Accumulate performs an operation on each element of the input
func Accumulate(input []string, converter func(string) string) []string {
	var output = make([]string, len(input))
	for i, str := range input {
		output[i] = converter(str)
	}
	return output
}