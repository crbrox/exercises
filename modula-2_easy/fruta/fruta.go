package main

import "fmt"

type FRUTA uint

const (
	manzana FRUTA = iota
	naranja
	plátano
	ciruela
	uva
)

var (
	TipoDeFruta FRUTA
	s           string
)

func main() {
	for {
		fmt.Print("introducir fruta: ")
		fmt.Scanln(&s)
		switch s[0] {
		case 'm':
			TipoDeFruta = manzana
		case 'n':
			TipoDeFruta = naranja
		case 'p':
			TipoDeFruta = plátano
		case 'c':
			TipoDeFruta = ciruela
		case 'u':
			TipoDeFruta = uva
		}
		fmt.Println()
		if TipoDeFruta == manzana {
			fmt.Println("Hoy están buenas las manzanas")
		} else if TipoDeFruta == naranja {
			fmt.Println("No hay naranjas hasta el martes")
		} else if TipoDeFruta == plátano {
			fmt.Println("Hay mucha cantidad")
		} else if TipoDeFruta == ciruela {
			fmt.Println("fuera de temporada")
		} else if TipoDeFruta == uva {
			fmt.Println("Decidir precio de venta")
		} else {
			fmt.Println("Intentálo de nuevo") // No debería escribirlo si es 'q', por fidelidad al original
		}
		if s[0] == 'q' {
			break
		}
		// Si introducimos un valor incorrecto, se queda asignado el mensaje anterior
		// "error" en la lógica del original ...
	}
}
