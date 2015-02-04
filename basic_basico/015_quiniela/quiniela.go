package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var (
		resultado string
		p         float64
	)
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 15; i++ {
		p = rand.Float64()
		switch {
		case 0 <= p && p < .5:
			resultado = "1"
		case .5 <= p && p < .5+.3:
			resultado = "X"
		default:
			resultado = "2"
		}
		fmt.Printf("PARTIDO %2d - SIGNO %s\n", i, resultado)
	}
}
