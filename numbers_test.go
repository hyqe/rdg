package rdg

import (
	"testing"
)

func TestUint63n(t *testing.T) {
	var maxs = []uint64{
		18446744073709551615,
		18446744073709551614,
		18446744073709551613,
		18446744073709551612,
		18446744073709551611,
		18446744073709551610,
		9223372036854775810,
		9223372036854775809,
		9223372036854775808,
		9223372036854775807,
		9223372036854775806,
		9223372036854775805,
		9223372036854775804,
		9223372036854775803,
		9223372036854775802,
		9223372036854775801,
		4294967296,
		4294967295,
		4294967294,
		3,
		2,
		1,
	}
	t.Parallel()
	t.Run("table", func(t *testing.T) {
		for _, max := range maxs {
			n := Uint63n(max)
			if n >= max {
				t.Fatalf("value is greater than max: %v", n)
			}
		}
	})
	t.Run("table_min_max", func(t *testing.T) {
		for _, max := range maxs {
			min := max - 1
			n := Uint63n(max-min) + min
			if n < min || n >= max {
				t.Fatalf("value %v was le than %v or ge than %v", n, min, max)
			}
		}
	})
	t.Run("low", func(t *testing.T) {
		var max uint64 = 1
		n := Uint63n(max)
		if n >= max {
			t.Fatalf("value is greater than max: %v", n)
		}
	})
	t.Run("high", func(t *testing.T) {
		var max uint64 = maxUint64
		n := Uint63n(max)
		if n >= max {
			t.Fatalf("value is greater than max: %v", n)
		}
	})
	t.Run("zero_panic", func(t *testing.T) {
		defer func(t *testing.T) {
			if r := recover(); r == nil {
				t.Fatal("failed to panic on zero")
			}
		}(t)
		Uint63n(0)
	})
}

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
	t.Run("eq_zero", func(t *testing.T) {
		var min int32 = -2147
		var max int32 = 2147
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
	t.Run("negative", func(t *testing.T) {
		var min int64 = -5
		var max int64 = -1
		i := Int64Between(min, max)
		if i < min || i > max {
			t.Fatalf("want between: %v and %v got: %v", min, max, i)
		}
	})
	t.Run("eq_zero", func(t *testing.T) {
		var min int64 = -5
		var max int64 = 5
		i := Int64Between(min, max)
		if i < min || i > max {
			t.Fatalf("want between: %v and %v got: %v", min, max, i)
		}
	})
}

func TestUint64Between(t *testing.T) {
	var numbers = []uint64{
		18446744073709551615,
		18446744073709551614,
		18446744073709551613,
		18446744073709551612,
		18446744073709551611,
		18446744073709551610,
		18446744073709551607,
		18446744073709551606,
		18446744073709551605,
		18446744073709551604,
		18446744073709551603,
		18446744073709551602,
		18446744073709551601,
		9223372036854775810,
		9223372036854775809,
		9223372036854775808,
		9223372036854775807,
		9223372036854775806,
		9223372036854775805,
		9223372036854775804,
		9223372036854775803,
		9223372036854775802,
		9223372036854775801,
		4294967296,
		4294967295,
		4294967294,
		3,
		2,
		1,
	}
	t.Parallel()
	t.Run("table", func(t *testing.T) {
		for _, max := range numbers {
			min := max - 1
			i := Uint64Between(min, max)
			if i < min || i > max {
				t.Fatalf("want between: %v and %v got: %v", min, max, i)
			}
		}
	})
	t.Run("equal", func(t *testing.T) {
		i := Uint64Between(1, 1)
		if i != 1 {
			t.Fatalf("want: %v got: %v", 1, i)
		}
	})
	t.Run("one_option", func(t *testing.T) {
		i := Uint64Between(1, 2)
		if i != 1 {
			t.Fatalf("want: %v got: %v", 1, i)
		}
	})
	t.Run("overflow", func(t *testing.T) {
		// TODO: define overflow behavor
	})
}
