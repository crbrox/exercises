// angulos
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/crbrox/exercises/basic"
)

func main() {
	// Conversión de ángulos
	// (Borrar pantalla)
	var Z int
	basic.Input("De grados a radianes (1) o al revés (2) ", &Z)
	if Z == 2 {
		rad2grad()
	} else {
		grad2rad()
	}
}

func grad2rad() {
	var linea string
	fmt.Println("Dar los grados, minutos y segundos")
	basic.Input("separados por comas (,) ", &linea)
	campos := strings.SplitN(linea, ",", 3)
	G, _ := strconv.ParseFloat(campos[0], 64)
	M, _ := strconv.ParseFloat(campos[1], 64)
	S, _ := strconv.ParseFloat(campos[2], 64)
	G = G + M/60 + S/3600
	R := G * math.Pi / 180
	fmt.Printf("%.3f radianes\n", R)
}
func rad2grad() {
	var R float64
	basic.Input("angulo en radianes ", &R)
	G := R * 180 / math.Pi
	G1 := int(G)
	G -= float64(G1) // G1, G := math.Modf(G)
	M := G * 60
	M1 := int(M)
	S := (M - float64(M1)) * 60
	S1 := int(S*1000+.5) / 1000
	fmt.Printf("%d grados %d minutos %d segundos\n", G1, M1, S1)
}
