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

func Test(t *testing.T) {
	cases := []testCase{
		{0, 0},
		{1, 1},
		{9, 1},
		{10, 2},
		{18446744073709551615, 21},
	}

	for _, c := range cases {
		actual := Log10(c.input)
		if actual != c.expected {
			t.Errorf("expected Log10(%d) = %d, but got %d", c.input, c.expected, actual)
		}
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
