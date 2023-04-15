// sieve eratosthenes

package main

import (
	"fmt"
)

const n = 10e6

func main() {

	/*
		f, err := os.Create("./cpuprof")
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	*/

	a := make([]bool, n)

	a[0] = true // is non prime
	a[1] = true // we start in 2
	for i, isNonPrime := range a {
		if !isNonPrime { // is prime
			for j := i; j*i < n; j++ {
				a[j*i] = true // is non-prime
			}
		}
	}

	count := 0
	for _, isNonPrime := range a {
		if !isNonPrime { // is prime
			count++
		}
	}
	fmt.Printf("%4d primes ", count)

}
