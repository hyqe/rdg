package rdg

import (
	"github.com/LukeOchoa/usefulThings"
)



func FirstNameFemale() string {
	return femaleNames[usefulThings.NewInt32Between(0, int32(len(femaleNames)))]
}
