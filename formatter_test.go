package rdg

import "testing"

var formatter = []Formatter{
	func(in string) string {
		return in + "1"
	},
	func(in string) string {
		return in + "2"
	},
	func(in string) string {
		return in + "3"
	},
}

func TestFormat(t *testing.T) {
	in := "ABC"
	for i := 0; i < 10; i++ {
		result := Format(in, formatter...)
		if result != (in+"1") && result != (in+"2") && result != (in+"3") {
			t.Fatalf("failed to format: %v", result)
		}
	}
}
