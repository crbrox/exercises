// servidor
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	var mensaje = make([]byte, 1024)
	fmt.Println("Hello World!")
	conn, err := net.ListenPacket("udp4", "localhost:2341")
	if err != nil {
		log.Fatal(err)
	}
	for {
		n, addr, err := conn.ReadFrom(mensaje)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v (%d) -> %s\n", addr, n, mensaje[:n])
	}

}
