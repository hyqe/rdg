package rdg

import "testing"

func TestFormat(t *testing.T) {
	formatter := []Formatter{
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

	t.Parallel()
	t.Run("table", func(t *testing.T) {
		in := "ABC"
		for i := 0; i < 10; i++ {
			result := Format(in, formatter...)
			if result != (in+"1") && result != (in+"2") && result != (in+"3") {
				t.Fatalf("failed to format: %v", result)
			}
		}
	})
	t.Run("nil", func(t *testing.T) {
		in := "ABC"
		x := Format(in)
		if x != in {
			t.Fatalf(`want: "%v" got: "%v"`, in, x)
		}
	})

}
