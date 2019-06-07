package main

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
	return nil, nil
}

func (h *handlers) broadcastHandler(connection *connection, payload interface{}) (interface{}, error) {
	return nil, nil
}
