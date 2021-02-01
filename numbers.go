package rdg

import (
	"math/rand"
)

const maxInt64 = 9223372036854775807
const minInt64 = -9223372036854775808
const maxUint64 = 18446744073709551615

// Uint63n generates a random number between 0 and (max-1).
// it behaves like rand.Int63n except it works with uint64.
// it will panic if max is equal to 0. this is to imitate the
// behvoior of the math/rand package.
func Uint63n(max uint64) uint64 {
	if max == 0 {
		panic("max can not be zero")
	}
	if max <= maxInt64 {
		return uint64(rand.Int63n(int64(max)))
	}

	// this is an exception case, where the largest
	// uint64 number is an ood number which can not
	// be stored and converted as a positive int64.
	if max == maxUint64 {
		// generate with the largest number possible.
		x := rand.Int63n(maxInt64)

		// maxUint64 1 larger then maxInt64 when stored.
		// so randomaly add 1 to enable the possibilty of
		// generateing a truely large and rangdom uint64.
		if Bool() {
			return unstoreInt64AsUint64(x) + 1
		}
		return unstoreInt64AsUint64(x)
	}

	a := storeUint64AsInt64(max)

	isNeg := false
	if a < 0 {
		isNeg = true
		a = a * -1 // to positive
	}

	b := rand.Int63n(a)

	if isNeg {
		b = b * -1 // back to neg
	}

	return unstoreInt64AsUint64(b)
}

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

// Uint64Between generates a random int between min and max.
// is min is greater then max the resulting uint64 will overflow.
func Uint64Between(min, max uint64) uint64 {
	if min == max {
		return min
	}
	return Uint63n(max-min) + min
}
