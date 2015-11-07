package main

import (
	"bytes"
	"log"
	"os"
	"syscall"
)

const bufSize = 10
const svSockPath = "/tmp/testgounixdgram"

func main() {
	log.Println("Hello World!, unix socket dgram server (hard)")

	// Borra path si ya existe
	os.Remove(svSockPath)

	// Creamos el server socket
	sfd, err := syscall.Socket(syscall.AF_UNIX, syscall.SOCK_DGRAM, 0)
	if err != nil {
		log.Fatalln("socket", err)
	}

	// Construimos la well-known direccion y ligamos el socket
	srvAddr := syscall.SockaddrUnix{Name: svSockPath}
	if err = syscall.Bind(sfd, &srvAddr); err != nil {
		log.Fatal(err)
	}

	// Recibir mensajes, convertirlos a mayúsculas y devolverlos al cliente
	for {
		msg := make([]byte, bufSize)
		size, sender, err := syscall.Recvfrom(sfd, msg, 0)
		if err != nil || size == -1 {
			log.Fatalln("recvfrom", err)
		}
		log.Printf("recibido mensaje %q (%d) de %v\n", msg[:size], size, pretty(sender))

		resp := bytes.ToUpper(msg[:size])

		err = syscall.Sendto(sfd, resp, 0, sender)
		if err != nil {
			log.Fatalln("sendto", err)
		}
	}
}

func pretty(addr syscall.Sockaddr) string {
	uAddr := addr.(*syscall.SockaddrUnix)
	return uAddr.Name
}
