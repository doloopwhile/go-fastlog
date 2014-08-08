package fastlog

// Log returns the decimal logarithm of n.
// Log10(0) = 0 as a special case.
//
// Log10 must very faster than math.Log10 in the standard library.
func Log10(n uint64) int {
	switch {
	case n == 0:
		return 0
	case n < 1e1:
		return 1
	case n < 1e2:
		return 2
	case n < 1e3:
		return 3
	case n < 1e4:
		return 4
	case n < 1e5:
		return 5
	case n < 1e6:
		return 6
	case n < 1e7:
		return 7
	case n < 1e8:
		return 8
	case n < 1e9:
		return 9
	case n < 1e10:
		return 10
	case n < 1e11:
		return 11
	case n < 1e12:
		return 12
	case n < 1e13:
		return 13
	case n < 1e14:
		return 14
	case n < 1e15:
		return 15
	case n < 1e16:
		return 16
	case n < 1e17:
		return 17
	case n < 1e18:
		return 18
	case n < 1e19:
		return 19
	default:
		return 20
	}
}
