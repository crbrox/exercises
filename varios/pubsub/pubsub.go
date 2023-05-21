package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
	"time"
)

type Subscriber[T any] struct {
	name   string
	events chan T
}

type PubSub[T any] struct {
	subscribers []Subscriber[T]
}

func NewPubSub[T any]() PubSub[T] {
	return PubSub[T]{
		subscribers: make([]Subscriber[T], 0, 10),
	}
}
func (ps *PubSub[T]) Subscribe(s Subscriber[T]) {
	if s.events == nil {
		panic("events channel is nil for " + s.name)
	}
	ps.subscribers = append(ps.subscribers, s)
}

func (ps *PubSub[T]) Publish(event T) {
	for _, s := range ps.subscribers {
		log.Printf("(pusub) sending event '%v' to %s\n", event, s.name)
		if s.events == nil {
			panic("events channel is nil for " + s.name)
		}
		// drop message if subscriber is not ready (could be a timeout instead)
		select {
		case s.events <- event:
		default:
		}
	}
}

var subscriberBufferSize uint

func main() {

	log.SetFlags(log.Lmicroseconds)

	flag.UintVar(&subscriberBufferSize, "size", 1, "size channel of the subscriber")
	flag.Parse()

	pubsub := NewPubSub[string]()
	wg := sync.WaitGroup{}
	stop := make(chan struct{})
	go_consumer("A", stop, &wg, &pubsub)
	go_consumer("B", stop, &wg, &pubsub)
	go_consumer("C", stop, &wg, &pubsub)

	log.Println("first round")
	pubsub.Publish("start #1")
	pubsub.Publish("processing #1")
	pubsub.Publish("end #1")

	time.Sleep(1 * time.Second)
	fmt.Println()
	log.Println("second round")

	pubsub.Publish("start #2")
	pubsub.Publish("processing #2")
	pubsub.Publish("end #2")

	time.Sleep(500 * time.Millisecond)

	close(stop)
	wg.Wait()

}

func go_consumer(name string, stop chan struct{}, wg *sync.WaitGroup, pubsub *PubSub[string]) {
	subs := Subscriber[string]{
		name:   name,
		events: make(chan string, subscriberBufferSize),
	}
	pubsub.Subscribe(subs)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-stop:
				return
			case event := <-subs.events:
				log.Println("consumer", name, "received event '", event, "'")
			}
		}
	}()
}
