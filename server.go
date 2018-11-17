package main

import (
	"log"
	"net/http"
)

const  (
	port = ":4200"
	dir = "./"
)

func main() {
	log.Printf("listening on %q...", port)
	log.Fatal(http.ListenAndServe(port, http.FileServer(http.Dir(dir))))
}