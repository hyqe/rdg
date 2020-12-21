package rdg

import "math/rand"

// Bool generates a random boolean.
func Bool() bool {
	return rand.Intn(2) == 1
}
