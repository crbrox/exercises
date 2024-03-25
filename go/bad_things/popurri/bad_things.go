package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {

	fmt.Println("El comienzo de una serie de catastróficas desdichas")

	mutatingIterate()

	notUsedErr()

	dataRace()

	// invalid use of mutex, which cannot be copied
	// w := createMutex()
	// lockMutex(w)
	// unlockMutex(w)

	fmt.Println("El fin de una serie de catastróficas desdichas")
}

type WrapMutex struct {
	m sync.Mutex
}

func createMutex() WrapMutex {
	return WrapMutex{sync.Mutex{}}
}

func lockMutex(w WrapMutex) {
	w.m.Lock()
}

func unlockMutex(w WrapMutex) {
	w.m.Unlock()
}

func mutatingIterate() {
	m := map[int]string{
		1: "uno",
		2: "dos",
		3: "tres",
	}
	for k, v := range m {
		m[k*3] = v + " por tres"
	}
	fmt.Printf("map: %v\n", m)
}

func notUsedErr() error {
	f, err := os.Open("non_existent_file.com")
	_, err = io.ReadAll(f)
	err = f.Close()
	return err
}

func dataRace() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	shared := []int{0}
	go func() {
		for i := 0; i < 7; i++ {
			shared = append(shared, i)
		}
		wg.Done()
	}()
	go func() {
		for i := 1; i < 8; i *= 2 {
			shared = append(shared, i)
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Printf("shared: %v\n", shared)
}

type A struct {
	Value int
}

func ValueFieldModified() {
	a := A{Value: 1}
	a.Change()
	fmt.Printf("a.Value = %d\n", a.Value)
}
func (a A) Change() {
	a.Value = 2
}
