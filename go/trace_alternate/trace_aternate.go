package main

import (
	"log"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

func main() {

	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	if err := trace.Start(f); err != nil {
		log.Fatalln(err)
	}
	defer trace.Stop()

	wg := &sync.WaitGroup{}

	{
		chSender := make(chan struct{})
		wg.Add(1)
		go sender(wg, chSender)
		wg.Add(1)
		go receiver(wg, chSender)
	}

	{
		time.Sleep(500 * time.Microsecond)
		chSignal := make(chan struct{})
		wg.Add(1)
		go signal(wg, chSignal)
		for i := 0; i < 3; i++ {
			wg.Add(1)
			go wait(wg, chSignal)
		}
	}

	wg.Wait()
}

func sender(wg *sync.WaitGroup, ch chan<- struct{}) {
	defer wg.Done()
	time.Sleep(10 * time.Microsecond)
	ch <- struct{}{}
}

func receiver(wg *sync.WaitGroup, ch <-chan struct{}) {
	defer wg.Done()
	<-ch
}

func signal(wg *sync.WaitGroup, ch chan<- struct{}) {
	defer wg.Done()
	time.Sleep(10 * time.Microsecond)
	close(ch)
}

func wait(wg *sync.WaitGroup, ch <-chan struct{}) {
	defer wg.Done()
	<-ch
}
