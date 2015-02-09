// moneda.go
package main

import (
	"fmt"
	"math/rand"

	"github.com/crbrox/exercises/basic"
)

func main() {
	var (
		k = 0
		n int
	)
	basic.Input("NUMERO DE LANZAMIENTOS ", &n)
	for i := 1; i <= n; i++ {
		a := rand.Intn(2)
		if a == 0 {
			fmt.Print("C")
			k++
		} else {
			fmt.Print("+")
		}
	}
	for i := 1; i < 5; i++ {
		fmt.Println()
	}
	fmt.Println("NUMERO DE LANZAMIENTOS", n)
	fmt.Println("NUMERO DE CARAS", k)
	fmt.Println("NUMERO DE CRUCES", n-k)
	fmt.Println("FRECUENCIA CARAS", float64(k)/float64(n))
	fmt.Println("FRECUENCIA CRUCES", float64(n-k)/float64(n))
}
