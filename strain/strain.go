package strain

// Ints is the type for a slice of ints
type Ints []int

// Lists is the type for a slice of a slice of ints
type Lists [][]int

// Strings is the type for a slice of ints
type Strings []string

// Keep returns a new Ints with elements from ints for which f evaluates to true
func (ints Ints) Keep(f func(int) bool) Ints {
	var kept Ints
	if len(ints) == 0 {
		return nil
	}

	for _, i := range ints {
		if result := f(i); result == true {
			kept = append(kept, i)
		}
	}
	return kept
}

// Discard returns a new Ints with elements from ints for which f evaluates to false
func (ints Ints) Discard(f func(int) bool) Ints {
	var kept Ints
	if len(ints) == 0 {
		return nil
	}

	for _, i := range ints {
		if result := f(i); result != true {
			kept = append(kept, i)
		}
	}
	return kept
}

// Keep returns a new String with elements from strings for which f evaluates to true
func (strings Strings) Keep(f func(string) bool) Strings {
	var kept Strings
	if len(strings) == 0 {
		return nil
	}

	for _, i := range strings {
		if result := f(i); result == true {
			kept = append(kept, i)
		}
	}
	return kept
}

// Keep returns a new [][]int with elements from Lists for which f evaluates to true
func(lists Lists) Keep(f func([]int) bool) Lists {
	if len(lists) == 0 {
		return nil
	}

	var kept Lists

	for _, list := range lists {
		if result := f(list); result == true {
			kept = append(kept, list)
		}
	}

	return kept
}