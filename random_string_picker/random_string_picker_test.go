package random_string_picker

import (
	"fmt"
	"testing"
)

var Tests = [][]string{
	{"one", "two", "three", "four"},
	{"Does", "this", "work?"},
	{"Is", "this", "string", "out", "of", "order"},
	{"100_100", "2ksjd99__", "009asdf3jJJA_)(&$$", "ju21j4165nasd+f8K"},
}

func ifStringIn(stringSlice []string, s string) bool {

	var boolean bool = false
	for _, value := range stringSlice {
		if s == value {
			boolean = true
		}
	}
	if !boolean {
		return false
	} else {
		return true
	}
}

func TestRandomStringPicker(t *testing.T) {
	for _, test := range Tests {
		value, err := random_string_picker(test...)
		if err != nil {
			t.Error("no words or slice/array of words was passed")
			t.Fatal(err)
		}
		if !ifStringIn(test, value) {
			fmt.Println("Failed to return a random string from strings passed.")
			t.FailNow()
		}
	}
}
