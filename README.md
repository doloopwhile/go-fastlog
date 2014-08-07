go-fastlog
==========

Simple and fast Log10 implementation for uint64.

It might be useful to calculate width of number as decimal.

# func Log10
```go
func Log10(n uint64) int
```
Log returns the decimal logarithm of n.
Log10(0) = 0 as a special case.

Log10 must very faster than math.Log10 in the standard library.
