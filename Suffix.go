package rdg

import (
	"github.com/LukeOchoa/usefulThings"
)

var suffixes = []string{
	"test",
	"data",
}

func Suffix() string {
	return suffixes[usefulThings.NewInt32Between(0, int32(len(suffixes)))]
}
