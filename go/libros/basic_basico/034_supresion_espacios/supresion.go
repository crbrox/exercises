//Haz un programa que suprima todos los espacios
//en blanco de una cadena
package main

import (
	"fmt"
	"github.com/crbrox/exercises/basic"
)

func main() {
	var cadena string
	basic.Input("ESCRIBE LA CADENA ", &cadena)
	for _, c := range cadena {
		if c != ' ' {
			fmt.Printf("%c", c)
		}
	}
	fmt.Println()
}
