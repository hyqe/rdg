package rdg

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"fmt"
)

var phoneRegex = regexp.MustCompile(`\+1[0-9]\d{9}`)
var extractNumbers = regexp.MustCompile(`\d+`)

func ParsePhone(phone string) (string, error) {

	phone = strings.Join(extractNumbers.FindAllString(phone, -1), "")

	if !validatePhone(phone) {
		return phone, errors.New("bad phone number")
	}

	return phone, nil
}

func validatePhone(number string) bool {

	boolean := phoneRegex.MatchString("+" + number)
	if !boolean {
		return false
	}
	return true
}

func stringNumber(digits, min, max int) string {
	var value string = ""
	for i := 0; i < digits; i++ {
		value = value + strconv.Itoa(IntBetween(min, max))
	}
	return value
}

func TestPhone(t *testing.T) {
	_, err := ParsePhone(PhoneNumber())
	if err != nil {
		t.Fatalf("PhoneNumber() failed to return a proper phone number")
	}

	var testCase CustomNumber
	for i := 0; i < 100; i++ {
		if i < 30 {
			testCase = CustomNumber{
				stringNumber(3, 0, 10),
				"0",
				"0",
			}
		}
		if i >= 30 && i < 60 {
			testCase = CustomNumber{
				"0",
				stringNumber(3, 0, 10),
				"0",
			}
		}
		if i >= 60 && i < 100 {
			testCase = CustomNumber{
				"0",
				"0",
				//stringNumber(4, 0, 10),
				"3333",
			}
		}
		value := CustomPhoneNumber(testCase)
		_, err2 := ParsePhone(value)
		if err2 != nil {
			t.Fatalf("CustomPhoneNumber() failed to return a proper phone number")
		}
		fmt.Printf("Phone Number: [ %v ]\n", value)
	}

}
