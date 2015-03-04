package main

import (
	"fmt"
	"math/rand"
)

type Arbol struct {
	Derecho   *Arbol
	Valor     int
	Izquierdo *Arbol
}

func recorre(a *Arbol, ch chan int) {
	if a == nil {
		return
	}
	recorre(a.Izquierdo, ch)
	ch <- a.Valor
	recorre(a.Derecho, ch)
}
func Recorre(a *Arbol, ch chan int) {
	recorre(a, ch)
	close(ch)
}

func Iguales(a1, a2 *Arbol) bool {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	go Recorre(a1, ch1)
	go Recorre(a2, ch2)

	for i1 := range ch1 {
		i2, ok := <-ch2
		if !ok || i1 != i2 {
			return false
		}
	}
	_, ok := <-ch2
	if ok { // Todavia quedan en t2
		return false
	}
	return true
}

func main() {
	ch := make(chan int)
	a1 := Nuevo(1)
	fmt.Println(a1)
	go Recorre(a1, ch)
	for i := range ch {
		fmt.Print(i, " -> ")
	}
	fmt.Println()
	a2 := Nuevo(1)
	fmt.Println(a2)
	ch = make(chan int)
	go Recorre(a2, ch)
	for i := range ch {
		fmt.Print(i, " -> ")
	}
	fmt.Println()
	fmt.Println("Iguales(a1,a2)", Iguales(a1, a2))
	a3 := Nuevo(2)
	fmt.Println(a3)
	ch = make(chan int)
	go Recorre(a3, ch)
	for i := range ch {
		fmt.Print(i, " -> ")
	}
	fmt.Println()
	fmt.Println("Iguales(a1,a3)", Iguales(a1, a3))
	fmt.Println("Iguales(a1,a1)", Iguales(a1, a1))
	fmt.Println("Iguales(a2,a2)", Iguales(a2, a2))
	fmt.Println("Iguales(a3,a3)", Iguales(a3, a3))
	fmt.Println("Iguales(nil,nil)", Iguales(nil, nil))
	fmt.Println("Iguales(&Arbol{},&Arbol{})", Iguales(&Arbol{}, &Arbol{}))
	fmt.Println("Iguales(&Arbol{Valor:1},&Arbol{Valor:1})", Iguales(&Arbol{Valor: 1}, &Arbol{Valor: 1}))
	fmt.Println("Iguales(&Arbol{},nil)", Iguales(&Arbol{}, nil))

	a4 := &Arbol{&Arbol{nil, 3, nil}, 4, &Arbol{nil, 5, nil}}
	a5 := &Arbol{nil, 3, &Arbol{nil, 4, &Arbol{nil, 5, nil}}}
	fmt.Println("Iguales(a4,a5)", Iguales(a4, a5))

}
func Inserta(a *Arbol, i int) *Arbol {
	if a == nil {
		return &Arbol{Valor: i}
	}
	if i < a.Valor {
		a.Izquierdo = Inserta(a.Izquierdo, i)
	} else {
		a.Derecho = Inserta(a.Derecho, i)
	}
	return a
}
func (a *Arbol) String() string {
	if a == nil {
		return " "
	}
	if a.Derecho == nil && a.Izquierdo == nil { // Estamos en una hoja}
		return fmt.Sprintf("(%d)", a.Valor)
	}
	return fmt.Sprintf("%d -> [%s,%s]", a.Valor, a.Izquierdo, a.Derecho)
}

func Nuevo(k int) *Arbol {
	valores := rand.Perm(10)
	for i := range valores {
		valores[i] *= k
	}
	var arbol *Arbol
	for _, valor := range valores {
		arbol = Inserta(arbol, valor)
	}
	return arbol
}
