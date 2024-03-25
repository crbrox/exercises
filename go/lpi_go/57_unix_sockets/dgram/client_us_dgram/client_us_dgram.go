package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const bufSize = 10 * 2
const svSockPath = "/tmp/testgounixdgram"
const clSockPathFormat = "/tmp/testgounixdgram-%d"

func main() {
	log.Println("Hello World!, unix socket dgram client")
	clPath := fmt.Sprintf(clSockPathFormat, os.Getpid())
	localAddr := &net.UnixAddr{Net: "unixgram", Name: clPath}
	serverAddr := &net.UnixAddr{Net: "unixgram", Name: svSockPath}
	//Esto es un simple bind, para poder recibir las respuestas
	unixConn, err := net.ListenUnixgram("unixgram", localAddr)
	// Hay que borrar el path al salir !
	defer os.Remove(clPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, msg := range os.Args[1:] {
		nw, err := unixConn.WriteToUnix([]byte(msg), serverAddr)
		log.Printf("(error: %v) enviados %d bytes\n", err, nw)
		if err != nil {
			// No esperamos repuesta si el envio no ha ido bien
			continue
		}
		resp := make([]byte, bufSize)
		size, sender, err := unixConn.ReadFromUnix(resp)
		log.Printf("(error:%v) recibido mensaje %q (%d) de %v\n", err, resp[:size], size, sender)
	}
}
