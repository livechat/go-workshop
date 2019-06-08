package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

var (
	host, prt = "localhost", 8085
)

// run chat server and wait for incoming websocket chatters
func init() { go startServer(host, prt) }

func TestIntegration_Login(t *testing.T) {
	type login struct{ Name, Avatar string }

	// set few websocket connection, emulating multiple persons visiting
	// chat window.
	c1 := newClient(host, prt, t)
	c2 := newClient(host, prt, t)
	c3 := newClient(host, prt, t)
	c4 := newClient(host, prt, t)
	c5 := newClient(host, prt, t)

	// let's join to the chat! people logs in to the server
	// in almost same time - let's see if we hit race condition.
	c1.send(login{"Tom", "tom.png"})
	c2.send(login{"Greg", "greg.png"})
	c3.send(login{"Kate", "kate.png"})
	c4.send(login{"Jimbo", "jimbo.jpg"})
	c5.send(login{"Joanna", "joanna.jpg"})

	time.Sleep(time.Second)
}

type client struct {
	c *websocket.Conn
	t *testing.T
}

func newClient(host string, port int, t *testing.T) *client {
	addr := fmt.Sprintf("ws://%s:%d/ws", host, port)

	c, _, err := (&websocket.Dialer{}).Dial(addr, nil)
	if err != nil {
		t.Fatal(err)
	}

	return &client{c, t}
}

func (s *client) send(payload interface{}) {
	type message struct {
		Action  string      `json:"action"`
		Payload interface{} `json:"payload"`
	}

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
