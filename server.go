package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type server struct {
	connections *connections
	handlers    *handlers
}

func newServer() *server {
	s := &server{
		connections: newConnections(),
	}
	s.handlers = newHandlers(s.connections)
	return s
}

func (s *server) websocketHandler(rw http.ResponseWriter, req *http.Request) {
	ws, err := websocket.Upgrade(rw, req, rw.Header(), 10240, 10240)
	if err != nil {
		log.Print(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	newConnection(ws, s.handlers, s.connections)
}
