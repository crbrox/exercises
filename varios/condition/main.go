package main

import (
	"flag"
	"log"
	"sync"
	"time"
)

var isOnline = true
var isOnlineMutex = &sync.Mutex{}

var signal bool

func main() {
	flag.BoolVar(&signal, "signal", false, "send two signals instead of a broadcast")
	flag.Parse()

	log.SetFlags(log.Ltime | log.Lmicroseconds)

	wg := &sync.WaitGroup{}

	// define a condition variable that monitors the server state
	serverStatus := sync.NewCond(isOnlineMutex)

	wg.Add(1)
	go connection("A", serverStatus, wg)

	wg.Add(1)
	go connection("B", serverStatus, wg)

	time.Sleep(1 * time.Second)

	log.Println("take the server offline")
	isOnlineMutex.Lock()
	isOnline = false
	isOnlineMutex.Unlock()

	time.Sleep(1 * time.Second)

	log.Println("take the server online")
	isOnlineMutex.Lock()
	isOnline = true
	isOnlineMutex.Unlock()

	if signal {
		// signal one, sleep enough and signal again
		log.Println("sending first signal")
		serverStatus.Signal()
		time.Sleep(3 * time.Second)
		log.Println("sending second signal")
		serverStatus.Signal()
	} else {
		// Notify all
		log.Println("sending a broadcast")
		serverStatus.Broadcast()
	}

	log.Println("main, wait goroutines")
	wg.Wait()

}
func connection(name string, serverStatus *sync.Cond, wg *sync.WaitGroup) {
	defer wg.Done()

	// read a message from the caller
	for x := 0; x < 30; x++ {
		// check the server is alive
		serverStatus.L.Lock()
		for !isOnline {
			log.Println(name, "blocked")
			serverStatus.Wait()
		}
		serverStatus.L.Unlock()
		sendToServer(name, x)
	}
}

func sendToServer(name string, x int) {
	log.Println(name, "sending ", x)
	time.Sleep(100 * time.Millisecond)

}
