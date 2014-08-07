package fastlog

// Log returns the decimal logarithm of n.
// Log10(0) = 0 as a special case.
//
// Log10 must very faster than math.Log10 in the standard library.
func Log10(n uint64) int {
	switch {
	case n == 0:
		return 0
	case n < 10:
		return 1
	case n < 100:
		return 2
	case n < 100:
		return 3
	case n < 1000:
		return 4
	case n < 10000:
		return 5
	case n < 100000:
		return 6
	case n < 1000000:
		return 7
	case n < 10000000:
		return 8
	case n < 100000000:
		return 9
	case n < 1000000000:
		return 10
	case n < 10000000000:
		return 11
	case n < 100000000000:
		return 12
	case n < 1000000000000:
		return 13
	case n < 10000000000000:
		return 14
	case n < 100000000000000:
		return 15
	case n < 1000000000000000:
		return 16
	case n < 10000000000000000:
		return 17
	case n < 100000000000000000:
		return 18
	case n < 1000000000000000000:
		return 19
	case n < 10000000000000000000:
		return 20
	default:
		return 21
	}
}
