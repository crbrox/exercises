package main

import "fmt"

func main() {
	var (
		x float64
	)
	fmt.Println()
	x = 1.0E-291 / 3.0
	for i := 1; i < 18; i++ {
		fmt.Printf("%.20E", x)
		x *= 0.1
		fmt.Printf("    %.20E\n", x)
		x *= 0.1
	}
}
