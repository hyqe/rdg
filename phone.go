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

func stringNumber(digits, min, max int) string {
        var value string = ""
        for i := 0; i < digits; i++ {
                value = value + strconv.Itoa(IntBetween(min, max))
        }
        return value
}

func CustomPhoneNumber(number CustomNumber) string {

	if len(number.area_code) < 3 {
		number.area_code = strconv.Itoa(IntBetween(0, 2)) + stringNumber(2, 0, 10)
	}
	if len(number.prefix) < 3 {
		number.prefix = stringNumber(3, 0, 10)
	}
	if len(number.line_number) < 4 {
		number.line_number = stringNumber(4, 0, 10)
	}

	customPhoneNumber := "1" + number.area_code + number.prefix + number.line_number

	return customPhoneNumber
}
