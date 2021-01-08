package rdg

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"
)

var regex = regexp.MustCompile(`^[-]?[$]\d+.\d\d$`)

func TestMoney(t *testing.T) {

	var min int64 = 0
	var minDec int64 = 0

	var max int64 = 9999
	var maxDec int64 = 99

	for i := 0; i < 100; i++ {
		ival := Money()
		if !regex.MatchString(ival) {
			t.Fatalf("Money Function failed, ( %v )", ival)
		}

		value := MoneyBetween(min, minDec, max, maxDec)
		if !regex.MatchString(value) {
			t.Fatalf("MoneyBetween() Function failed, ( %v )", value)
		}

		var tempValue string
		for xval := 0; xval < len(value); xval++ {
			if x := string([]rune(value)[xval]); x != "$" {
				tempValue = tempValue + x
			}
		}

		valueFloat, err := strconv.ParseFloat(tempValue, 64)
		if err != nil {
			panic(err)
		}

		realMin, err2 := strconv.ParseFloat(fmt.Sprintf("%v.%v", min, minDec), 64)
		if err2 != nil {
			panic(err2)
		}
		realMax, err3 := strconv.ParseFloat(fmt.Sprintf("%v.%v", max, maxDec), 64)
		if err3 != nil {
			panic(err3)
		}
		if valueFloat < realMin || valueFloat > realMax {
			t.Fatalf("Value returned from MoneyBetween() was out of range")
		}
	}
}
