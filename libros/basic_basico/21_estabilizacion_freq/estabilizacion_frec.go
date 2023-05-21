/*
Vamos a simular el lanzamiento de una moneda, el programa parará
cuando se estabilicen las frecuencias relativas. Después de cada
lanzamiento se halla la frecuencia relativa de las caras. Cuando
esta frecuencia difiera de la del lanzamiento anterior en menos
de 0.01 el programa parará, inidicando el número de lanzamientos
el de caras y el de cruces, ylas dos últimas frecuencias relativas
de caras redondeadas a 4 decimales.
*/
package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	var n, c, k int
	var fr [2]float64
	for {
		for i := k; i < 2; i++ {
			a := rand.Intn(2)
			n++
			if a == 0 {
				c++
				fmt.Print("C")
			} else {
				fmt.Print("X")
			}
			fr[i] = float64(c) / float64(n)
		}
		d := math.Abs(fr[0] - fr[1])
		if d < 0.01 && fr[0] != fr[1] {
			break
		}
		fr[0] = fr[1]
		k = 1 // mejorable, pero así en el original
	}
	fmt.Println()
	fmt.Println("Hemos efectuado ", n, " lanzamientos")
	fmt.Println("Hemos obtenido ", c, "caras y ", n-c, " cruces")
	fmt.Printf("Las dos últimas frecuencias son %.4f y %.4f\n", fr[0], fr[1])
}
