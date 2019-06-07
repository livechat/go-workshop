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

func init() {
	go startServer(host, prt)
}

func TestIntegration_Login(t *testing.T) {

	c1 := newClient(host, prt, t)
	c2 := newClient(host, prt, t)
	c3 := newClient(host, prt, t)
	c4 := newClient(host, prt, t)
	c5 := newClient(host, prt, t)

	type login struct{ Name, Avatar string }

	go c1.send(login{"Tom", "tom.png"})
	go c2.send(login{"Greg", "greg.png"})
	go c3.send(login{"Kate", "kate.png"})
	go c4.send(login{"Jimbo", "jimbo.jpg"})
	go c5.send(login{"Joanna", "joanna.jpg"})
	go c5.send(login{"JoannaX", "joanna.jpg"})
	go c5.send(login{"JoannaY", "joanna.jpg"})
	go c5.send(login{"JoannaZ", "joanna.jpg"})

	time.Sleep(time.Second)

}

func TestIntegration_Message(t *testing.T) {
	var (
		mike = newClient(host, prt, t)
		greg = newClient(host, prt, t)
	)

	type (
		login     struct{ Name, Avatar string }
		broadcast struct{ Text string }
	)

	// conversation begins!
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

	c, _, err := websocket.DefaultDialer.Dial(addr, nil)
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
