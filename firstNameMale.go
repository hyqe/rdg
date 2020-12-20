package rdg

import (
	"github.com/LukeOchoa/usefulThings"
)



func FirstNameMale() string {
	return maleNames[usefulThings.NewInt32Between(0, int32(len(maleNames)))]
}
