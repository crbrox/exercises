package main

import (
	"fmt"
	"math"
)

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
	fmt.Println("Arcsin")
	fmt.Printf("%-20.20s | %-20.20s | %-20.20s\n", "x", "arcsin(x)", "x")
	for i := 0; i < 6; i++ {
		x := float64(i) * 0.2
		y := math.Asin(x)
		z := math.Sin(y)
		fmt.Printf("%20.16f | %20.16f | %20.16f\n", x, y, z)
	}
	fmt.Println("Raiz cuadrada")
	fmt.Printf("%-20.20s | %-20.20s | %-20.20s\n", "x", "sqrt(x)", "x")
	for i := 0; i < 6; i++ {
		x := float64(i) * 0.2
		y := math.Sqrt(x)
		z := math.Pow(y, 2)
		fmt.Printf("%20.16f | %20.15f | %20.16f\n", x, y, z)
	}
	fmt.Println("Exp")
	fmt.Printf("%-20.20s | %-20.20s | %-20.20s\n", "x", "exp(x)", "x")
	for i := 0; i < 6; i++ {
		x := float64(i) * 0.2
		y := math.Exp(x)
		z := math.Log(y)
		fmt.Printf("%20.16f | %20.16f | %20.16f\n", x, y, z)
	}
}
