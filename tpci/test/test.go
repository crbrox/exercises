package main

import (
	"fmt"
	"math"
)

func prueba(nombre string, funcion, inversa func(float64) float64) {
	fmt.Println(nombre)
	fmt.Printf("%-20.20s | %-20.20s | %-20.20s\n", "x", nombre+"(x)", "x")
	for i := 0; i < 6; i++ {
		x := float64(i) * 0.2
		y := funcion(x)
		z := inversa(y)
		fmt.Printf("%20.16f | %20.16f | %20.16f\n", x, y, z)
	}
}

func main() {
	var (
		x float64
	)
	fmt.Println("Precisión float64")
	x = 1.0E-291 / 3.0
	for i := 1; i < 18; i++ {
		fmt.Printf("%.16E", x)
		x *= 0.1
		fmt.Printf("    %.16E\n", x)
		x *= 0.1
	}
	prueba("Arcsin", math.Asin, math.Sin)
	prueba("Sqrt", math.Sqrt, func(x float64) float64 { return math.Pow(x, 2) })
	prueba("Exp", math.Exp, math.Log)
}
