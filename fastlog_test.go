package fastlog

import (
	"fmt"
	"math"
	"testing"
)

type testCase struct {
	input    uint64
	expected int
}

func assertLog10(t *testing.T, n uint64) {
	actual := Log10(n)
	expected := len(fmt.Sprint(n))
	if actual != expected {
		t.Errorf("expected Log10(%d) = %d, but got %d", n, expected, actual)
	}
}

func Test(t *testing.T) {
	cases := []uint64{
		10,
		1e3,
		1e4,
		1e5,
		1e6,
		1e7,
		1e8,
		1e9,
		1e10,
		1e11,
		1e12,
		1e13,
		1e14,
		1e15,
		1e16,
		1e17,
		1e18,
		1e19,
		18446744073709551615,
	}

	assertLog10(t, 1)

	for _, x := range cases {
		for _, n := range []uint64{x, x - 1, x / 2} {
			assertLog10(t, n)
		}
	}

	actual := Log10(0)
	if actual != 0 {
		t.Errorf("expected Log10(0) = 0, but got %d", actual)
	}
}

const benchInput = 54321

func BenchmarkFastLog10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Log10(54321)
	}
}

func BenchmarkMathLog10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = math.Log10(54321.0)
	}
}

func BenchmarkFmtPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = len(fmt.Sprint(54321))
	}
}
