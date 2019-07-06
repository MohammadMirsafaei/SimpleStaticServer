package main

import (
	"net/http"
	"log"
	"flag"
	"strconv"
)
func main() {
	// Parsing flags
	port := flag.Int("port",8080,"port number for server")
	dir := flag.String("dir","static","directory to serve")

	// Handling server
	fs := http.FileServer(http.Dir(*dir))
	http.Handle("/",fs)

	// Logging
	log.Println("Static file server listening on: ",*port)

	// Listening
	http.ListenAndServe(":"+strconv.Itoa(*port),nil)

}