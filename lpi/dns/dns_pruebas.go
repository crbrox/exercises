package main

import (
	"log"
	"net"
	"os"
)

func main() {
	for _, h := range os.Args[1:] {
		addrs, err := net.LookupHost(h)
		log.Println(addrs, err)
		ips, err := net.LookupIP(h)
		log.Println(ips, err)
	}

}
