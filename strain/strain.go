package strain

// Ints is the type for a slice of ints
type Ints []int

// Lists is the type for a slice of a slice of ints
type Lists [][]int

// Strings is the type for a slice of ints
type Strings []string

// Keep returns a new Ints with elements from ints for which f evaluates to true
func (ints Ints) Keep(f func(int) bool) Ints {
	if len(ints) == 0 {
		return nil
	}

	

}