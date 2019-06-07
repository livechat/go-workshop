package main

type connections struct {
	list map[string]*connection
}

func newConnections() *connections {
	return &connections{
		list: make(map[string]*connection),
	}
}

func (c *connections) register(name string, connection *connection) error {
	return nil
}

func (c *connections) unregister(name string) error {
	return nil
}

func (c *connections) broadcastText(author, text string) {
}

func (c *connections) broadcastUsersDetails() {
}

func (c *connections) broadcast(msg interface{}) {
	for _, connection := range c.list {
		connection.writeC <- msg
	}
}
