package main

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type connection struct {
	id          string
	name        string
	avatar      string
	socket      *websocket.Conn
	handlers    *handlers
	connections *connections
	writeC      chan interface{}
}

func newConnection(socket *websocket.Conn, handlers *handlers, connections *connections) *connection {
	c := &connection{
		socket:      socket,
		handlers:    handlers,
		connections: connections,
		writeC:      make(chan interface{}),
	}

	go c.reader()
	go c.writer()

	log.Print("new connection")

	return c
}

func (c *connection) reader() {
}

func (c *connection) writer() {
}

func (c *connection) handleRequest(req *request) {
}

func (c *connection) sendSuccess(action string, payload interface{}) {
	response := &response{
		Action:  action,
		Success: true,
		Payload: payload,
	}

	c.writeC <- response
}

func (c *connection) sendError(action, msg string) {
	response := &response{
		Action: action,
		Error:  msg,
	}

	c.writeC <- response
}

func decode(src []byte, dst interface{}) error {
	return json.Unmarshal(src, dst)
}
