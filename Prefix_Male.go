package rdg

import (
	"github.com/LukeOchoa/usefulThings"
)

var malePrefixes = []string{
	"test",
	"data",
}

func PrefixMale() string {
	return malePrefixes[usefulThings.NewInt32Between(0, int32(len(malePrefixes)))]
}
