package aleatorio

import (
	"math/rand"
	"testing"
)

func BenchmarkFloat32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Float32()
	}
}
func BenchmarkFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Float64()
	}
}
func BenchmarkInt31(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int31()
	}
}
func BenchmarkInt63(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int31()
	}
}
func BenchmarkUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Uint32()
	}
}
func BenchmarkExpFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.ExpFloat64()
	}
}
func BenchmarkNormFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.NormFloat64()
	}
}
