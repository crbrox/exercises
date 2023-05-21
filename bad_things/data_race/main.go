package main

import (
	"fmt"
	"sync"
)

func Example() {
	inputs := []int{1, 2, 3, 4, 5}
	outputs := make([]int, len(inputs))

	wg := &sync.WaitGroup{}
	wg.Add(len(inputs))

	for index, value := range inputs { // <- Previous write by main goroutine
		go func() {
			defer wg.Done()
			outputs[index] = value * 2 // <- Read  by goroutine 8:
		}()
	}
	wg.Wait()
	for index, value := range outputs {
		fmt.Println(index, "->", value)
	}
}
func main() {
	Example()
}
