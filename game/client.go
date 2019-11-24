package game

import (
	"bufio"
	"fmt"
	"net"
)

// Client is a state manager for the client
type Client struct {
	// State stores State logic
	State *State
	// Name is the username
	Name string
	// Conn holds the connection to the server
	Conn net.Conn
}

// Listen listens for messages
func (c *Client) Listen() {
	reader := bufio.NewReader(c.Conn)

	for {
		data, err := reader.ReadString('\n')

		if err == nil {
			switch t, payload := Parse(data); t {
			case ChangeName:
				fmt.Println(payload)

			}
		}
	}
}

// ConfirmUsername sets a username and informs the server about it
func (c *Client) ConfirmUsername(username string) {
	c.Name = username
	c.Conn.Write(Compose(ChangeName, username))
}
