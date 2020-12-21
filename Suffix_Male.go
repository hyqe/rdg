package rdg

import (
	"github.com/LukeOchoa/usefulThings"
)

var maleSuffixes = []string{
	"test",
	"data",
}

func SuffixMale() string {
	return maleSuffixes[usefulThings.NewInt32Between(0, int32(len(maleSuffixes)))]
}
