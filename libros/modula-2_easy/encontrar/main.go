package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	Frase [80]byte
	i     int
	ch    byte
	in    = bufio.NewReader(os.Stdin)
)

func main() {
	for i := 0; i < 79; i++ {
		Frase[i] = ' ' // inicializamos el array
	}
	fmt.Print("Introduzca una frase: ")
	Frase, _ := in.ReadBytes('\n') //Uuups, descartando error, por respeto al original
	fmt.Print("Introduzca el caracter a localizar: ")
	ch, _ = in.ReadByte()
	i = 0
	for Frase[i] != ch && i < 80 {
		fmt.Printf("%c", Frase[i])
		i++
	}
	fmt.Println()
	fmt.Printf("Primer %c está en el lugar %3d\n", ch, i)
}
