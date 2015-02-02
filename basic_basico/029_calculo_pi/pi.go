package main

import (
	"fmt"
	"math/rand"

	"github.com/crbrox/exercises/basic"
)

func main() {
	m := 0
	n := 0
	basic.Input("¿Cuántos dardos lanzas? ", &n)
	for i := 0; i < n; i++ {
		x := rand.Float64()
		y := rand.Float64()
		if x*x+y*y < 1 {
			m++
		}
	}
	// (Borrar pantalla)
	fmt.Printf("Si lanzas %d dardos obtenemos\n", n)
	fmt.Printf("un valor aproximado de Pi = %f\n", 4*float64(m)/float64(n))
}
