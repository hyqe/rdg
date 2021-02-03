package rdg

import (
	"errors"
	"math/rand"
	"time"
)

func random_string_picker(list ...string) (string, error) {

	var anError error
	if len(list) < 0 {
		anError = errors.New("Only one or zero items in list")
	} else {
		anError = nil
	}
	rand.Seed(time.Now().UnixNano())
	randomValue := rand.Intn(len(list))

	return list[randomValue], anError
}
