package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
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
			fmt.Print("Escribe tu numero ")
			fmt.Scanln(&n)
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
			var err error
			fmt.Print("¿Quieres volver a jugar (S/N)?")
			resp, err = input.ReadString('\n')
			if err != nil {
				os.Exit(1)
			}
			resp = resp[:len(resp)-1]
		}
		otra = resp == "S"

	}
	fmt.Println("adios")
}
