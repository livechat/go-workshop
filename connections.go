package main

import (
	"fmt"
	"log"
)

type connections struct {
	list     map[string]*connection
	messages uint64
}

func newConnections() *connections {
	return &connections{
		list: make(map[string]*connection),
	}
}

func (c *connections) register(name string, connection *connection) error {
	if _, exists := c.list[name]; exists {
		return fmt.Errorf("connection %s already registered", name)
	}

	c.list[name] = connection

	log.Printf("connection registered: %s", name)

	return nil
}

func (c *connections) unregister(name string) error {
	if _, exists := c.list[name]; !exists {
		return fmt.Errorf("connection %s does not exist", name)
	}

	delete(c.list, name)

	log.Printf("connection unregistered: %s", name)

	return nil
}

func (c *connections) broadcastText(author, text string) {
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
		connection.writeC <- msg
	}
}
