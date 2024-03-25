package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

const delayInSeconds = 2

type SlowListener struct {
	net.Listener
}

func (sl SlowListener) Accept() (c net.Conn, err error) {
	log.Println("Conexión entrante")
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

	server := &http.Server{}
	server.SetKeepAlivesEnabled(false)
	log.Fatal(server.Serve(SlowListener{lsnr}))
}
