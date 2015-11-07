package main

import (
	"log"
	"os"
	"syscall"
)

const bufSize = 100
const svSockPath = "/tmp/testgounixsocket"

func main() {
	log.Println("Hello World!, unix socket stream client (hard)")

	// create socket
	sockfd, err := syscall.Socket(syscall.AF_UNIX, syscall.SOCK_STREAM, 0)
	if err != nil {
		log.Fatalln("socket", err)
	}

	// connect to server
	err = syscall.Connect(sockfd, &syscall.SockaddrUnix{Name: svSockPath})
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, bufSize)
	for nr, readErr := os.Stdin.Read(data); readErr == nil; nr, readErr = os.Stdin.Read(data) {
		log.Printf("leidos %d bytes, error: %v\n", nr, readErr)

		// Write to Unix socket
		nw, writeErr := syscall.Write(sockfd, data[:nr])

		log.Printf("escritos %d/%d bytes, error: %v\n", nw, nr, writeErr)

	}
}
