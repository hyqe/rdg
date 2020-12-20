package rdg

import (
	"fmt"
	"math/rand"
)

// IntBetween returns a random int between min and max.
// i := IntBetween(min,max) would be (i >= min && i <= max)
func IntBetween(min, max int) int {
	if min > max {
		panic(fmt.Sprintf("min can not be greater then max. min: %v, max: %v", min, max))
	}
	if min == max {
		return min
	}
	return rand.Intn(max-min) + min
}
