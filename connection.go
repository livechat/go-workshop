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
		c.handleRequest(req)
	}
}

func (c *connection) sendMessage(msg interface{}) {
	if err := c.socket.WriteJSON(msg); err != nil {
		log.Printf("error sending message: %v", err)
	}
}

func (c *connection) handleRequest(req *request) {
	log.Print("request: ", req.Action)

	if len(req.RawPayload) == 0 {
		log.Printf("error: empty payload")
		c.sendError(req.Action, "empty payload")
		return
	}

	switch req.Action {
	case "login":
		req.Payload = &requestLogin{}
	case "broadcast":
		req.Payload = &requestBroadcast{}
	default:
		log.Printf("error: unknown action: %s", req.Action)
		c.sendError(req.Action, "unknown action")
		return
	}

	if err := decode(req.RawPayload, req.Payload); err != nil {
		log.Printf("error unmarshaling payload: %v", err)
		c.sendError(req.Action, "internal")
		return
	}

	resPayload, err := c.handlers.list[req.Action](c, req.Payload)
	if err != nil {
		log.Printf("error handling request: %v", err)
		c.sendError(req.Action, err.Error())
		return
	}

	c.sendSuccess(req.Action, resPayload)
}

func (c *connection) sendSuccess(action string, payload interface{}) {
	response := &response{
		Action:  action,
		Success: true,
		Payload: payload,
	}

	c.sendMessage(response)
}

func (c *connection) sendError(action, msg string) {
	response := &response{
		Action: action,
		Error:  msg,
	}

	c.sendMessage(response)
}

func decode(src []byte, dst interface{}) error {
	return json.Unmarshal(src, dst)
}