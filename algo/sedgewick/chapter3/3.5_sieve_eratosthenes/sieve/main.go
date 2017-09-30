// sieve eratosthenes

package main

import (
	"bufio"
	"fmt"
	"os"
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

	output := bufio.NewWriter(os.Stdout)

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
	for i, isNonPrime := range a {
		if !isNonPrime { // is prime
			fmt.Fprintf(output, "%4d ", i)
		}
	}
	fmt.Println()
}
func fatalError(x interface{}) {
	fmt.Println(x)
	os.Exit(-1)
}
