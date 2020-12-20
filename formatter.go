package rdg

import (
	"math/rand"
)

// Formatter is any function that can format a string
// Formatters should never panic.
type Formatter = func(string) string

// Format randomly formats a string with a list of formatters
func Format(in string, fns ...Formatter) string {
	if len(fns) <= 0 {
		return in
	}
	index := rand.Intn(len(fns) - 1)
	return fns[index](in)
}
