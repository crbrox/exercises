package main

import (
	"fmt"
	"math/rand"
	"time"
)

// LA CARTA LLEGARA EL LUNES
const NUM_PALABRAS = 5
const N = 20

var estado [NUM_PALABRAS]int

var palabras = [NUM_PALABRAS][]string{
	{"LA"},
	{"CARTA", "TARTA", "MARTA", "SARTA"},
	{"LLEGARA", "LLOVERA", "LLAMARA"},
	{"EL", "DEL"},
	{"LUNES", "MARTES"},
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < N; i++ {
		// cambiamos el estado
		for j := 0; j < NUM_PALABRAS; j++ {
			if rand.Float32() > 0.93 {
				estado[j] = (estado[j] + 1) % len(palabras[j])
			}
			fmt.Print(palabras[j][estado[j]], " ")
		}
		fmt.Println()

	}
}
