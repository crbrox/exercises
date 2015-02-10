package main

import (
	"fmt"

	"github.com/crbrox/exercises/basic"
)

func main() {
	const Max = 80

	var xs []float64
	var media, dest float64
	var N int

	fmt.Println()
	fmt.Println("\tCálculo de la media y la desviación estándar")
	for {
		basic.Input("\t¿Cuántos puntos? ", &N)
		if N <= Max {
			break
		}
	}
	xs = make([]float64, N)
	for i := 0; i < N; i++ {
		basic.Input(fmt.Sprintf("%3d: ", i+1), &xs[i])
	}
	media, dest = medEst(xs)
	fmt.Println()
	fmt.Printf("\tPara %3d puntos Media = %8.4f Sigma = %8.4f\n", N, media, dest)
}
