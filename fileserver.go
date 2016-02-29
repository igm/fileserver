package main

import (
	"flag"
	"net/http"

	"log"
)

var listenAddr = flag.String("l", ":8080", "Listen port")

func init() {
	flag.Parse()
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	log.Println("Listening on", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}
