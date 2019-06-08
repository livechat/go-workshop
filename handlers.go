package main

import (
	"errors"
	"log"
)

type handlers struct {
	connections *connections
	list        map[string]handler
}

type handler func(connection *connection, payload interface{}) (interface{}, error)

func newHandlers(connections *connections) *handlers {
	h := &handlers{
		connections: connections,
	}
	h.list = map[string]handler{
		"login":     h.loginHandler,
		"broadcast": h.broadcastHandler,
	}
	return h
}

func (h *handlers) loginHandler(connection *connection, payload interface{}) (interface{}, error) {
	pl, ok := payload.(*requestLogin)
	if !ok {
		log.Print("error: invalid login payload")
		return nil, errors.New("internal")
	}

	if pl.Name == "" {
		return nil, errors.New("`name` is required")
	}

	if pl.Avatar == "" {
		return nil, errors.New("`avatar` is required")
	}

	connection.name = pl.Name
	connection.avatar = pl.Avatar

	name := pl.Name
	if err := h.connections.register(name, connection); err != nil {
		return nil, err
	}

	h.connections.broadcastUsersDetails()

	return nil, nil
}

func (h *handlers) broadcastHandler(connection *connection, payload interface{}) (interface{}, error) {
	pl, ok := payload.(*requestBroadcast)
	if !ok {
		log.Print("error: invalid broadcast payload")
		return nil, errors.New("internal")
	}

	if pl.Text == "" {
		return nil, errors.New("`text` is required")
	}

	h.connections.broadcastText(connection.name, pl.Text)

	return nil, nil
}
