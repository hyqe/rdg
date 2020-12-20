package rdg

import (
	"github.com/LukeOchoa/usefulThings"
)

var femalePrefixes = []string{
	"test",
	"data",
}

func PrefixFemale() string {
	return femalePrefixes[usefulThings.NewInt32Between(0, int32(len(femalePrefixes)))]
}
