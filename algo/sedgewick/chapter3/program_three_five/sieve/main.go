// sieve eratosthenes

package main

import (
	"fmt"
)

const n = 1000

var a [n]bool

func main() {
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
