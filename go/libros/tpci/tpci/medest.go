package tcpi

import (
	"math"
	"math/rand"
)

func MedEst(xs []float64) (media float64, devEst float64) {
	N := float64(len(xs))
	sumX := 0.0
	sumCdr := 0.0
	for _, x := range xs {
		sumX += x
		sumCdr += x * x
	}
	media = sumX / N
	devEst = math.Sqrt((sumCdr - sumX*sumX/N) / (N - 1))
	return media, devEst
}

// Genera números aleatorios con una distribución gaussiana
func Aleatg(media, sigma float64) float64 {
	var (
		sum float64
	)
	for i := 0; i < 12; i++ {
		sum += rand.Float64()
	}
	return (sum-6)*sigma + media
}
