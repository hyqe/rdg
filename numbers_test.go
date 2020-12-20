package rdg

import "testing"

func TestIntBetween(t *testing.T) {
	t.Parallel()
	t.Run("one_option", func(t *testing.T) {
		i := IntBetween(1, 1)
		if i != 1 {
			t.Fatalf("want: %v got: %v", 1, i)
		}
	})
	t.Run("two_option", func(t *testing.T) {
		min := 1
		max := 2
		i := IntBetween(min, max)
		if i < min || i > max {
			t.Fatalf("want between: %v - %v got: %v", min, max, i)
		}
	})
	t.Run("min_greater_than_max", func(t *testing.T) {
		defer func(t *testing.T) {
			if r := recover(); r == nil {
				t.Fatal("failed to panic")
			}
		}(t)
		IntBetween(2, 1)
	})
	t.Run("low", func(t *testing.T) {
		min := -2147483647
		max := -2147483646
		i := IntBetween(min, max)
		if i < min || i > max {
			t.Fatalf("want between: %v - %v got: %v", min, max, i)
		}
	})
	t.Run("mid", func(t *testing.T) {
		min := -1
		max := 0
		i := IntBetween(min, max)
		if i < min || i > max {
			t.Fatalf("want between: %v - %v got: %v", min, max, i)
		}
	})
	t.Run("high", func(t *testing.T) {
		min := 2147483646
		max := 2147483647
		i := IntBetween(min, max)
		if i < min || i > max {
			t.Fatalf("want between: %v - %v got: %v", min, max, i)
		}
	})
}
