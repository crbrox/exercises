// server_su_stream project main.go
package main

import (
	"io"
	"log"
	"os"
	"syscall"
)

const backlog = 5
const bufSize = 100
const svSockPath = "/tmp/testgounixsocket"

func main() {
	log.Println("Hello World!, unix socket stream server (hard)")
	// Borrar path anterior
	syscall.Unlink(svSockPath)

	// create listening socket
	sfd, err := syscall.Socket(syscall.AF_UNIX, syscall.SOCK_STREAM, 0)
	if err != nil {
		log.Fatalln("socket", err)
	}

	// bind server path
	if err := syscall.Bind(sfd, &syscall.SockaddrUnix{Name: svSockPath}); err != nil {
		log.Fatalln("bind", err)
	}

	// listen
	if err := syscall.Listen(sfd, backlog); err != nil {
		log.Fatalln("listen", err)
	}
	for {

		// connection socket
		cfd, _, err := syscall.Accept(sfd)
		if err != nil {
			log.Fatalln("accept", err)
		}
		data := make([]byte, bufSize)
		var readErr error
		for readErr == nil {

			nr, readErr := syscall.Read(cfd, data)

			log.Printf("leidos %d bytes, error %v\n", nr, readErr)
			if readErr != nil {
				break
			}
			if nr <= 0 {
				readErr = io.EOF
				break
			}
			nw, writetErr := os.Stdout.Write(data[:nr])
			log.Printf("escritos %d bytes, error %v\n", nw, writetErr)
		}

		err = syscall.Close(cfd)

		if err != nil {
			log.Fatalln("close", err)
		}
	}
}
