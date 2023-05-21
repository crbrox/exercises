package main

import (
	"fmt"
	"log"
	"sync"
)

const QueueSize = 10

type Op struct {
	a, b int
	res  chan<- int
}
type Calc struct {
	ops  chan Op
	stop chan struct{}
}

func NewCalc() Calc {
	return Calc{
		ops:  make(chan Op, QueueSize),
		stop: make(chan struct{}),
	}
}

func (c Calc) On() {
	go func() {
		for {
			select {
			case <-c.stop:
				log.Println("cerrando Calc")
			case op := <-c.ops:
				res := op.a + op.b
				op.res <- res
			}
		}
	}()
}

func (c Calc) Off() {
	c.stop <- struct{}{}
}

func (c Calc) Sum(a, b int) int {
	res := make(chan int, 1)
	c.ops <- Op{
		a:   a,
		b:   b,
		res: res,
	}
	return <-res
}

func main() {
	calc := NewCalc()
	calc.On()

	wg := &sync.WaitGroup{}
	for i := 0; i < 6; i++ {
		go_sum(calc, i+1, wg)
	}
	wg.Wait()
	calc.Off()
}

func go_sum(calc Calc, i int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		res := calc.Sum(i, i*i)
		fmt.Println(i, "-> ", res)
	}()
}
