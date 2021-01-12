package rdg

import (
	"errors"
)

func SelectAny(list ...interface{}) (interface{}, error) {
	var err error
	if len(list) < 0 {
		err = errors.New("Only one or zero items in list")
	} else {
		err = nil
	}

	return list[IntBetween(0, len(list))], err
}
