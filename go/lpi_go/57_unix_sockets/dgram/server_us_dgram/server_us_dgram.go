package main

import (
	"bytes"
	"log"
	"net"
	"os"
)

const bufSize = 10
const svSockPath = "/tmp/testgounixdgram"

func main() {
	log.Println("Hello World!, unix socket dgram server")
	// Borra path si ya existe
	os.Remove(svSockPath)
	//Esto es un simple bind
	unixConn, err := net.ListenUnixgram("unixgram", &net.UnixAddr{Net: "Unixgram", Name: svSockPath})
	if err != nil {
		log.Fatal(err)
	}
	for {
		msg := make([]byte, bufSize)
		sender := &net.UnixAddr{}
		size, sender, err := unixConn.ReadFromUnix(msg)
		log.Printf("(error:%v) recibido mensaje %q (%d) de %v\n", err, msg[:size], size, sender)
		if err != nil {
			// No enviamos repuesta si no hemos recibido el mensaje bien
			continue
		}
		resp := bytes.ToUpper(msg[:size])
		nw, err := unixConn.WriteToUnix(resp, sender)
		log.Printf("(error: %v) enivados %d bytes", err, nw)
	}
}
