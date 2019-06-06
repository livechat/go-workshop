package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8000

func main() {
	server := newServer()

	http.HandleFunc("/ws", server.websocketHandler)

	log.Printf("starting listening on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), nil))
}
