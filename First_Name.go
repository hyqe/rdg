package rdg

import (
	"github.com/LukeOchoa/usefulThings"
)



func FirstName() string {
	return firstNames[usefulThings.NewInt32Between(0, int32(len(firstNames)))]
}
