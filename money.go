package rdg

import (
	"fmt"
	"github.com/LukeOchoa/usefulThings"
)

func Money() string {

	var str string
	var special string = "$"
	value := usefulThings.NewInt32()
	value2 := usefulThings.NewInt32Between(0, 10)
	value3 := usefulThings.NewInt32Between(0, 10)
	if value < 0 {
		value = value * -1
		special = "-$"
	}
	str = fmt.Sprintf("%v%v.%v%v", special, value, value2, value3)

	return str

}

func MoneyBetween(min, minDec, max, maxDec int64) string {

	var special string = "$"
	value := usefulThings.NewInt64Between(min, max)
	valueDec := usefulThings.NewInt64Between(minDec, maxDec)
	if value < 0 {
		value = value * -1
		special = "-$"
	}

	var xvalueDec string
	if valueDec < 10 {
		xvalueDec = fmt.Sprintf("0%v", valueDec)
	} else {
		xvalueDec = fmt.Sprintf("%v", valueDec)
	}

	str := fmt.Sprintf("%v%v.%v", special, value, xvalueDec)

	return str
}
