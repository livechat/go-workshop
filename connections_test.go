package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"sync"
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

	// let's join to the chat, multiple people logs in to the server
	// in almost same time - let's see if we have race condition.
	newClient(host, prt, t).send(login{"Tom", "tom.png"})
	newClient(host, prt, t).send(login{"Greg", "greg.png"})
	newClient(host, prt, t).send(login{"Kate", "kate.png"})
	newClient(host, prt, t).send(login{"Jimbo", "jimbo.jpg"})
	newClient(host, prt, t).send(login{"Joanna", "joanna.jpg"})

	time.Sleep(time.Second * 5)
}

func TestIntegration_Message(t *testing.T) {
	type (
		login     struct{ Name, Avatar string }
		broadcast struct{ Text string }
	)

	var (
		mike = newClient(host, prt, t)
		greg = newClient(host, prt, t)
	)

	// conversation begins! let's check if we get race conditions!
	mike.send(login{"Mike", "mike.jpg"})
	mike.send(broadcast{"Is there anybody outh there?"})
	greg.send(login{"Tom", "tom.png"})
	mike.send(broadcast{"Hi, im new here ;>"})
	greg.send(broadcast{"Oh! really?"})

}

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

func (s *client) send(payload interface{}) *client {
	type message struct {
		Action  string      `json:"action"`
		Payload interface{} `json:"payload"`
	}

	s.m.Lock()
	defer s.m.Unlock()

	if err := s.c.WriteJSON(message{reflect.TypeOf(payload).Name(), payload}); err != nil {
		s.t.Fatal(err)
	}

	return s
}

func startServer(host string, port int) {
	http.HandleFunc("/ws", newServer().websocketHandler)

	addr := fmt.Sprintf("%s:%d", host, port)
	log.Printf("starting server %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
