package rdg

import (
	"math/rand"
	"time"
)

// init the seed for the entire package
func init() {

	// TODO: confirm that the global seed
	// is cared over to the other files.
	rand.Seed(time.Now().UnixNano())
}
