// Sabiendo que la solución es única, hallar un número
// de cuatro cifras de la forma aabb, que sea cuadrado perfecto
package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	x, pruebas := metodo1()
	fmt.Printf("método 1 -> %d (%d)\n", x, pruebas)
	x, pruebas = metodo2()
	fmt.Printf("método 2 -> %d (%d)\n", x, pruebas)
}

func metodo1() (x, pruebas int) {
	for i := 1; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			x = 1100*i + 11*j
			pruebas++
			if r := int(math.Sqrt(float64(x))); r*r == x { // Es un cuadrado perfecto
				return x, pruebas
			}
		}
	}
	panic("metodo1: Debería haber uno")
}
func metodo2() (x, pruebas int) {
	for i := int(math.Sqrt(1111)); i <= int(math.Sqrt(9999)); i++ {
		x = i * i
		cadena := strconv.Itoa(i * i)
		pruebas++
		if cadena[0] == cadena[1] && cadena[2] == cadena[3] { // Es de la forma aabb
			return x, pruebas
		}
	}
	panic("metodo2: Deberia haber uno")
}
