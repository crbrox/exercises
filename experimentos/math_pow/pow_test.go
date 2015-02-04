package math_pow

import (
	"math"
	"testing"
)

var (
	x  float64 = 10
	y  float64 = 18
	iy         = 18
)

func BenchmarkPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Pow(x, y)
	}
}
func BenchmarkFor(b *testing.B) {
	var res = x
	for i := 0; i < b.N; i++ {
		for j := 1; j < iy; j++ {
			res *= x
		}
	}
}
