package tcpi

import (
	"math"
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
