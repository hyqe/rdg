package rdg

import (
	"math/rand"
)

// IntBetween generates a random int between min and max.
// if min is greater then max the returned value will be
// negative.
func IntBetween(min, max int) int {
	if min == max {
		return min
	}
	if min > max {
		i := rand.Intn(min-max) + max
		return i * -1
	}
	return rand.Intn(max-min) + min
}

// Int32Between generates a random int between min and max.
// if min is greater then max the returned value will be
// negative.
func Int32Between(min, max int32) int32 {
	if min == max {
		return min
	}
	if min > max {
		i := rand.Int31n(min-max) + max
		return i * -1
	}
	return rand.Int31n(max-min) + min
}

// Int64Between generates a random int between min and max.
// if min is greater then max the returned value will be
// negative.
func Int64Between(min, max int64) int64 {
	if min == max {
		return min
	}
	if min > max {
		i := rand.Int63n(min-max) + max
		return i * -1
	}
	return rand.Int63n(max-min) + min
}
