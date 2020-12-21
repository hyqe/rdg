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
			t.Fatalf("want between: %v and %v got: %v", min, max, i)
		}
	})
	t.Run("min_greater_than_max", func(t *testing.T) {
		min := 1
		max := -1
		i := IntBetween(min, max)
		if i >= min && i <= max {
			t.Fatalf("want between: %v and %v got: %v", min, max, i)
		}
	})
	t.Run("low", func(t *testing.T) {
		min := -2147483647
		max := -2147483646
		i := IntBetween(min, max)
		if i < min || i > max {
			t.Fatalf("want between: %v and %v got: %v", min, max, i)
		}
	})
	t.Run("mid", func(t *testing.T) {
		min := -1
		max := 0
		i := IntBetween(min, max)
		if i < min || i > max {
			t.Fatalf("want between: %v and %v got: %v", min, max, i)
		}
	})
	t.Run("high", func(t *testing.T) {
		min := 2147483646
		max := 2147483647
		i := IntBetween(min, max)
		if i < min || i > max {
			t.Fatalf("want between: %v and %v got: %v", min, max, i)
		}
	})
}

func TestInt32Between(t *testing.T) {
	t.Parallel()
	t.Run("one_option", func(t *testing.T) {
		i := Int32Between(1, 1)
		if i != 1 {
			t.Fatalf("want: %v got: %v", 1, i)
		}
	})
	t.Run("two_option", func(t *testing.T) {
		var min int32 = 1
		var max int32 = 2
		i := Int32Between(min, max)
		if i < min || i > max {
			t.Fatalf("want between: %v and %v got: %v", min, max, i)
		}
	})
	t.Run("min_greater_than_max", func(t *testing.T) {
		var min int32 = 1
		var max int32 = -1
		i := Int32Between(min, max)
		if i >= min && i <= max {
			t.Fatalf("want between: %v and %v got: %v", min, max, i)
		}
	})
	t.Run("low", func(t *testing.T) {
		var min int32 = -2147483647
		var max int32 = -2147483646
		i := Int32Between(min, max)
		if i < min || i > max {
			t.Fatalf("want between: %v and %v got: %v", min, max, i)
		}
	})
	t.Run("mid", func(t *testing.T) {
		var min int32 = -1
		var max int32 = 0
		i := Int32Between(min, max)
		if i < min || i > max {
			t.Fatalf("want between: %v and %v got: %v", min, max, i)
		}
	})
	t.Run("high", func(t *testing.T) {
		var min int32 = 2147483646
		var max int32 = 2147483647
		i := Int32Between(min, max)
		if i < min || i > max {
			t.Fatalf("want between: %v and %v got: %v", min, max, i)
		}
	})
}

func TestInt64Between(t *testing.T) {
	t.Parallel()
	t.Run("one_option", func(t *testing.T) {
		i := Int64Between(1, 1)
		if i != 1 {
			t.Fatalf("want: %v got: %v", 1, i)
		}
	})
	t.Run("two_option", func(t *testing.T) {
		var min int64 = 1
		var max int64 = 2
		i := Int64Between(min, max)
		if i < min || i > max {
			t.Fatalf("want between: %v and %v got: %v", min, max, i)
		}
	})
	t.Run("min_greater_than_max", func(t *testing.T) {
		var min int64 = 1
		var max int64 = -1
		i := Int64Between(min, max)
		if i >= min && i <= max {
			t.Fatalf("want between: %v and %v got: %v", min, max, i)
		}
	})
	t.Run("low", func(t *testing.T) {
		var min int64 = -9223372036854775808
		var max int64 = -9223372036854775807
		i := Int64Between(min, max)
		if i < min || i > max {
			t.Fatalf("want between: %v and %v got: %v", min, max, i)
		}
	})
	t.Run("mid", func(t *testing.T) {
		var min int64 = -1
		var max int64 = 0
		i := Int64Between(min, max)
		if i < min || i > max {
			t.Fatalf("want between: %v and %v got: %v", min, max, i)
		}
	})
	t.Run("high", func(t *testing.T) {
		var min int64 = 9223372036854775806
		var max int64 = 9223372036854775807
		i := Int64Between(min, max)
		if i < min || i > max {
			t.Fatalf("want between: %v and %v got: %v", min, max, i)
		}
	})
}
