package main

import (
	"log"
	"net"
	"os"
)

const bufSize = 100
const svSockPath = "/tmp/testgounixsocket"

func main() {
	log.Println("Hello World!, unix socket stream client")

	// Connect to server
	conn, err := net.Dial("unix", svSockPath)

	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, bufSize)
	for nr, readErr := os.Stdin.Read(data); readErr == nil; nr, readErr = os.Stdin.Read(data) {
		log.Printf("leidos %d bytes, error: %v", nr, readErr)

		// Write to Unix socket
		nw, writeErr := conn.Write(data[:nr])

		log.Printf("escritos %d bytes", nw)
		if writeErr != nil {
			log.Println(writeErr)
		}
	}
}
