package rdg

import (
	"github.com/LukeOchoa/usefulThings"
)

var lastNames = []string{
	"test",
	"data",
}

func LastName() string {
	return lastNames[usefulThings.NewInt32Between(0, int32(len(lastNames)))]
}
