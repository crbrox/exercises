// server_su_stream project main.go
package main

import (
	"log"
	"net"
	"os"
)

const backlog = 5
const bufSize = 100
const svSockPath = "/tmp/testgounixsocket"

func main() {
	log.Println("Hello World!, unix socket stream server")
	// Borrar path anterior
	os.Remove(svSockPath)

	// listening socket
	ln, err := net.Listen("unix", svSockPath)

	if err != nil {
		log.Fatal(err)
	}

	for {

		// connection socket
		conn, err := ln.Accept()

		if err != nil {
			log.Fatal(err)
		}
		data := make([]byte, bufSize)
		var readErr error
		for readErr == nil {

			nr, readErr := conn.Read(data)

			log.Printf("leidos %d bytes, error %v\n", nr, readErr)
			if readErr != nil {
				break
			}
			nw, writetErr := os.Stdout.Write(data[:nr])
			log.Printf("escritos %d bytes, error %v\n", nw, writetErr)
		}

		err = conn.Close()

		if err != nil {
			log.Fatal(err)
		}
	}
}
