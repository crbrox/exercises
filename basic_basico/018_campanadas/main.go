// ¿Cuantas campanadas da un reloj desde la hora H hasta las 12 de la noche
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("uso: campanas <hora de inicio>")
		os.Exit(-1)
	}
	hora, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("hora debe ser un entero,", err)
		os.Exit(-1)
	}
	if hora < 0 || hora > 24 {
		fmt.Println("la hora debe estar entre 0 y 24,", hora)
		os.Exit(-1)
	}
	if hora == 0 {
		hora = 24
	}
	campanadas := 0
	for i := hora; i <= 24; i++ {
		j := i
		if i > 12 {
			j -= 12
		}
		campanadas += j
		fmt.Printf("%2d (%2d) %s\n", i, j, strings.Repeat("\U0001F514 ", j))
	}
	fmt.Printf("Total de campanadas desde %2d:00 hasta las 24:00, %d\n", hora, campanadas)
}
