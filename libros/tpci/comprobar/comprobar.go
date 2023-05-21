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
	M          = 48
	N          = 20 * M
)

func porTrozos(numbers [N]float64) {
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
func main() {
	const ()
	var numbers [N]float64

	fmt.Println("UNIFORME (math/rand)")
	// Generamos todos los números
	for i := range numbers {
		numbers[i] = rand.Float64()
	}
	porTrozos(numbers)

	fmt.Println("GAUSS (tcpi)")
	// Generamos todos los números con dist. gaussiana
	for i := range numbers {
		numbers[i] = tcpi.Aleatg(MEDIA, DESVIACION)
	}
	porTrozos(numbers)

	fmt.Println("GAUSS (math/rand)")
	// Generamos todos los números con dist. gaussiana
	for i := range numbers {
		numbers[i] = rand.NormFloat64()*DESVIACION + MEDIA
	}
	porTrozos(numbers)

}
