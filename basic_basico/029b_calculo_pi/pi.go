package main

import (
	"flag"
	"fmt"
	"math/rand"

	"runtime"
)

var (
	n = flag.Int("lanzamientos", 1e6, "número de lanzamientos")
	g = flag.Int("goroutinas", runtime.NumCPU()-1, "número de goroutinas")
)

func lanzamiento(n int) (m int) {
	var x, y float64
	for i := 0; i < n; i++ {
		x = rand.Float64()
		y = rand.Float64()
		if x*x+y*y < 1 {
			m++
		}
	}
	return m
}
func main() {

	//Lo primero es lo primero
	flag.Parse()

	// Todo lo paralelo que podamos, pero que pueda respirar
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)

	m := 0 // lanzamientos dentro de la circunferencia
	if *g == 1 {
		m = lanzamiento(*n)
	} else {
		parciales := make(chan int, *g)
		for j := 0; j < *g; j++ {
			go func() {
				parciales <- lanzamiento(*n / *g)
			}()
		}
		for j := 0; j < *g; j++ {
			m += <-parciales
		}
	}
	// (Borrar pantalla)
	fmt.Printf("Si lanzas %d dardos obtenemos con %d rutinas\n", *n, *g)
	fmt.Printf("un valor aproximado de Pi = %f\n", 4*float64(m)/float64(*n))
}
