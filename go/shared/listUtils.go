package shared

// return all unique values of a given list
func Set[T comparable](lst []T) []T {
	setList := []T{}
	exists := map[T]bool{}
	for _, val := range lst {
		if !exists[val] {
			exists[val] = true
			setList = append(setList, val)
		}
	}
	return setList
}
