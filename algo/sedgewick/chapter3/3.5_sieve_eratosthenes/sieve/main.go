// sieve eratosthenes

package main

import (
	"fmt"
	"os"
	"strconv"
)

var a []bool

func main() {
	if len(os.Args) != 2 {
		fatalError("Usage: " + os.Args[0] + " <max number to check>")
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fatalError(err)
	}
	if n < 2 {
		fatalError(fmt.Sprint("it should be greater than 2, passed in ", n))
	}
	a = make([]bool, n)
	a[0] = true // is non prime
	a[1] = true // we start in 2
	for i, isNonPrime := range a {
		if !isNonPrime { // is prime
			for j := i; j*i < n; j++ {
				a[j*i] = true // is non-prime
			}
		}
	}
	for i, isNonPrime := range a {
		if !isNonPrime { // is prime
			fmt.Printf("%4d ", i)
		}
	}
	fmt.Println()
}
func fatalError(x interface{}) {
	fmt.Println(x)
	os.Exit(-1)
}
