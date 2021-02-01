package rdg

// storeUint64InInt64 stores a uint64 in a int64.
// even numbers will be stored as positive numbers.
// odd numbers will be stored as negative numbers.
// zero will be stored as zero.
//
// convertion table:
//   uint64 | int64
//        1 | -1
//        2 |  1
//        3 | -2
//        4 |  2
//
func storeUint64AsInt64(x uint64) int64 {
	if x == 0 {
		return 0
	}
	if x%2 == 0 {
		return int64(x / 2)
	}
	return (int64(x/2) + 1) * -1
}

// unstoreInt64FromUint64 unstores a int64 back into a uint64
// this function is only useful for retrieving a uint from an
// int64 that was previously converted from a uint64.
//
// convertion table:
//   int64 | uint64
//      -1 | 1
//       1 | 2
//      -2 | 3
//       2 | 4
//
func unstoreInt64AsUint64(x int64) uint64 {
	if x == 0 {
		return 0
	}
	if x < 0 {
		return (uint64(x*-1) * 2) - 1
	}
	return uint64(x) * 2
}
