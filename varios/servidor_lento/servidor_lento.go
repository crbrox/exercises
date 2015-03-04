package main

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

var data = [][]byte{
	[]byte("primera linea\n"),
	[]byte("segunda linea\n"),
	[]byte("tercera linea\n")}
var dataLen = 0

func main() {
	for _, b := range data {
		dataLen += len(b)
	}
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		resp.Header().Set("content-type", "text/plain")
		resp.Header().Set("content-length", strconv.Itoa(dataLen))
		resp.WriteHeader(200)
		for _, b := range data {
			resp.Write(b)
			if f, ok := resp.(http.Flusher); ok {
				f.Flush()
			}
			time.Sleep(3 * time.Second)
		}
	})
	log.Fatal(http.ListenAndServe(":4321", nil))
}
