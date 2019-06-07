package main

import (
	"fmt"
	"log"
	"sync"
)

type connections struct {
	list     map[string]*connection
	mu       sync.Mutex
	messages uint64
}

func newConnections() *connections {
	return &connections{
		list: make(map[string]*connection),
	}
}

func (c *connections) register(name string, connection *connection) error {
	// NO MUTEX ON PURPOSE
	if _, exists := c.list[name]; exists {
		return fmt.Errorf("connection %s already registered", name)
	}

	c.list[name] = connection

	log.Printf("connection registered: %s", name)

	return nil
}

func (c *connections) broadcastText(author, text string) {
	// NO MUTEX ON PURPOSE

	push := &push{
		Action: "message",
		Payload: &pushMessage{
			Author: author,
			Text:   text,
		},
	}

	c.broadcast(push)
	c.messages++

	log.Printf("`%s` sent #%d message number", author, c.messages)
}

func (c *connections) broadcastUsersDetails() {
	// NO MUTEX ON PURPOSE
	usersDetails := make([]*userDetails, len(c.list))

	var i int
	for _, connection := range c.list {
		usersDetails[i] = &userDetails{
			Name:   connection.name,
			Avatar: connection.avatar,
		}
		i++
	}

	push := &push{
		Action: "users",
		Payload: &pushUsers{
			List: usersDetails,
		},
	}

	c.broadcast(push)
}

func (c *connections) broadcast(msg interface{}) {
	for _, connection := range c.list {
		connection.sendMessage(msg)
	}
}
