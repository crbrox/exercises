package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"

	"github.com/crbrox/exercises/basic"
)

var input = bufio.NewReader(os.Stdin)

func main() {

	fmt.Println(`Tienes que adivinar un número
comprendido entre 1 y 100.
Después de escribir tu número
pulsa la tecla RETURN.`)

	otra := true // Otra partida?
	for otra {
		acertado := false   // Ha conseguido acertar
		x := rand.Intn(101) // Número oculto
		i := 0              // Número de intentos
		for !acertado {
			var n int // Suposición del jugador
			basic.Input("Escribe tu numero ", &n)
			i++
			if n < x {
				fmt.Println("Más alto")
			} else if n > x {
				fmt.Println("Más bajo")
			} else {
				fmt.Println("Enhorabuena. Acertaste")
				fmt.Printf("Después de %d intentos\n", i)
				acertado = true
			}
		}
		var resp string
		for resp != "N" && resp != "S" {
			basic.Input("¿Quieres volver a jugar (S/N)?", &resp)
		}
		otra = resp == "S"

	}
	fmt.Println("adios")
}
