package rdg

import (
	"github.com/LukeOchoa/usefulThings"
)

var femaleSuffixes = []string{
	"test",
	"data",
}

func SuffixFemale() string {
	return femaleSuffixes[usefulThings.NewInt32Between(0, int32(len(femaleSuffixes)))]
}
