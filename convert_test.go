package rdg

import "testing"

func TestStoreUint64AsInt64(t *testing.T) {
	numbers := map[uint64]int64{
		0:                    0,
		1:                    -1,
		2:                    1,
		3:                    -2,
		4:                    2,
		5:                    -3,
		6:                    3,
		7:                    -4,
		8:                    4,
		18446744073709551613: -9223372036854775807,
		18446744073709551614: 9223372036854775807,
		18446744073709551615: -9223372036854775808,
	}
	for u, i := range numbers {
		x := storeUint64AsInt64(u)
		if x != i {
			t.Fatalf("want: %v got: %v", i, x)
		}
	}
}

func TestUnstoreInt64AsUint64(t *testing.T) {
	numbers := map[int64]uint64{
		0:                    0,
		-1:                   1,
		1:                    2,
		-2:                   3,
		2:                    4,
		-3:                   5,
		3:                    6,
		-4:                   7,
		4:                    8,
		-9223372036854775807: 18446744073709551613,
		9223372036854775807:  18446744073709551614,
		-9223372036854775808: 18446744073709551615,
	}
	for i, u := range numbers {
		x := unstoreInt64AsUint64(i)
		if x != u {
			t.Fatalf("want: %v got: %v", u, x)
		}
	}
}
