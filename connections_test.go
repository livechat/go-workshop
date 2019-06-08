package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"sync"
	"testing"

	"github.com/gorilla/websocket"
)

type client struct {
	c *websocket.Conn
	t *testing.T
	m sync.Mutex
}

func newClient(host string, port int, t *testing.T) *client {
	addr := fmt.Sprintf("ws://%s:%d/ws", host, port)

	c, _, err := (&websocket.Dialer{}).Dial(addr, nil)
	if err != nil {
		t.Fatal(err)
	}

	return &client{c, t, sync.Mutex{}}
}

func (s *client) send(payload interface{}) {
	type message struct {
		Action  string      `json:"action"`
		Payload interface{} `json:"payload"`
	}

	s.m.Lock()
	defer s.m.Unlock()

	if err := s.c.WriteJSON(message{reflect.TypeOf(payload).Name(), payload}); err != nil {
		s.t.Fatal(err)
	}
}

func startServer(host string, port int) {
	http.HandleFunc("/ws", newServer().websocketHandler)

	addr := fmt.Sprintf("%s:%d", host, port)
	log.Printf("starting server %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
