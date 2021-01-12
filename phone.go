package rdg

import (
	"strconv"
)

func PhoneNumber() string {
	var number string = "1" + strconv.Itoa(IntBetween(0, 2))
	var min int = 0
	var max int = 10
	for i := 0; i < 9; i++ {
		number = number + strconv.Itoa(IntBetween(min, max))
	}

	return number
}

type CustomNumber struct {
	area_code   string
	prefix      string
	line_number string
}

func CustomPhoneNumber(number CustomNumber) string {

	var value string = ""
	if len(number.area_code) < 3 {
		for i := 0; i < 2; i++ {
			value = value + strconv.Itoa(IntBetween(0, 10))
		}
		number.area_code = strconv.Itoa(IntBetween(0, 2)) + value
		value = ""
	}
	if len(number.prefix) < 3 {
		for i := 0; i < 3; i++ {
			value = value + strconv.Itoa(IntBetween(0, 10))
		}
		number.prefix = value
		value = ""
	}
	if len(number.line_number) < 4 {
		for i := 0; i < 4; i++ {
			value = value + strconv.Itoa(IntBetween(0, 10))
		}
		number.line_number = value
		value = ""
	}

	customPhoneNumber := "1" + number.area_code + number.prefix + number.line_number

	return customPhoneNumber
}
