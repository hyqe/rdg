package rdg

// Bool generates a random boolean.
func Bool() bool {
	i := IntBetween(1, 2)
	if i == 1 {
		return true
	}
	return false
}
