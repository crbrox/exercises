// cliente
package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("Hello World!")
	for _, arg := range os.Args[1:] {
		mensaje := []byte(arg)
		conn, err := net.Dial("udp", "localhost:2341")
		if err != nil {
			log.Fatal(err)
		}
		n, err := conn.Write(mensaje)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v (%d) -> %s\n", conn.LocalAddr(), n, mensaje)
	}
}
