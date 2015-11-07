package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

const bufSize = 10 * 2
const svSockPath = "/tmp/testgounixdgram"
const clSockPathFormat = "/tmp/testgounixdgram-%d"

func main() {
	log.Println("Hello World!, unix socket dgram client (hard)")

	// Creamos el socket, el asignamos un pathname unico basado en el PID
	sfd, err := syscall.Socket(syscall.AF_UNIX, syscall.SOCK_DGRAM, 0)
	if err != nil {
		log.Fatalln("socket", err)
	}
	clPath := fmt.Sprintf(clSockPathFormat, os.Getpid())
	clAddr := syscall.SockaddrUnix{Name: clPath}
	err = syscall.Bind(sfd, &clAddr)
	if err != nil {
		log.Fatalln("bind", err)
	}

	defer os.Remove(clPath)

	// Direccion del servidor
	serverAddr := syscall.SockaddrUnix{Name: svSockPath}

	for j, msg := range os.Args[1:] {
		err := syscall.Sendto(sfd, []byte(msg), 0, &serverAddr)
		if err != nil {
			log.Fatalln("sendto", err)
		}
		resp := make([]byte, bufSize)
		numBytes, sender, err := syscall.Recvfrom(sfd, resp, 0)
		if err != nil {
			log.Fatalln("recvfrom", err)
		}
		log.Printf("respuesta #%d %q (%d) de %v\n", j+1, resp[:numBytes], numBytes, pretty(sender))
	}
}

func pretty(addr syscall.Sockaddr) string {
	uAddr := addr.(*syscall.SockaddrUnix)
	return uAddr.Name
}
