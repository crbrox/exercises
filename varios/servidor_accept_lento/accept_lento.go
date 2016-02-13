package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

const delayInSeconds = 10

type SlowListener struct {
	net.Listener
}

func (sl SlowListener) Accept() (c net.Conn, err error) {
	time.Sleep(delayInSeconds * time.Second)
	return sl.Listener.Accept()
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "Everything is alrigth!")
	})

	lsnr, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.Serve(SlowListener{lsnr}, nil))
}
