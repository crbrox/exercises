// Verificacion de la "calidad" del generador
// de números aleatorios, comprobando la media y
// la desviación típica.
package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/crbrox/exercises/tpci/tpci"
)

const (
	MEDIA      = 0.5
	DESVIACION = 0.2887
)

func main() {
	const N = 10000
	const M = 1000
	var numbers [N]float64
	// Generamos todos los números
	for i := range numbers {
		numbers[i] = rand.Float64()
	}
	//sacamos la media y la desviación
	//en bloques de M elementos
	fmt.Printf("%-8s %8s\n", "Media", "Desviación")
	fmt.Printf("%2.6f %2.6f\n", MEDIA, DESVIACION)
	fmt.Println(strings.Repeat("=", 8), strings.Repeat("=", 8))
	for j := 0; j < N; j += M {
		m, d := tcpi.MedEst(numbers[j : j+M])
		fmt.Printf("%2.6f %2.6f (%+2.6f) (%+2.6f)\n", m, d, m-MEDIA, d-DESVIACION)
	}

}
