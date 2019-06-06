package main

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type connection struct {
	id          string
	name        string
	avatar      string
	socket      *websocket.Conn
	handlers    *handlers
	connections *connections
}

func newConnection(socket *websocket.Conn, handlers *handlers, connections *connections) *connection {
	c := &connection{
		socket:      socket,
		handlers:    handlers,
		connections: connections,
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go c.reader(wg)
	wg.Wait()

	log.Print("new connection")

	return c
}

func (c *connection) reader(wg *sync.WaitGroup) {
	wg.Done()
	for {
		req := &request{}
		if err := c.socket.ReadJSON(req); err != nil {
			log.Printf("disconnect: %v", err)
			c.socket.Close()
			// TODO: remove from connections
			return
		}
		c.messageHandler(req)
	}
}

func (c *connection) sendMessage(msg interface{}) {
	if err := c.socket.WriteJSON(msg); err != nil {
		log.Printf("error sending message: %v", err)
	}
}

func (c *connection) messageHandler(req *request) {
	log.Print("request: ", req.Action)

	if len(req.RawPayload) == 0 {
		log.Printf("error: empty payload")
		// TODO: send response with error
		return
	}

	switch req.Action {
	case "login":
		req.Payload = &requestLogin{}
	case "broadcast":
		req.Payload = &requestBroadcast{}
	default:
		log.Printf("error: unknown action: %s", req.Action)
		// TODO: send response with error
		return
	}

	if err := json.Unmarshal(req.RawPayload, req.Payload); err != nil {
		log.Printf("error unmarshaling payload: %v", err)
		// TODO: send response with error
		return
	}

	_, err := c.handlers.list[req.Action](c, req.Payload)
	if err != nil {
		log.Printf("error handling request: %v", err)
		// TODO: send response with error
		return
	}
	// TODO: send response
}
