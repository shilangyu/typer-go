package game

import (
	"bufio"
	"fmt"
	"net"
)

// Client is a state manager for the client
type Client struct {
	State
	Name string
	Conn net.Conn
}

// Listen listens for messages
func (c *Client) Listen() {
	reader := bufio.NewReader(c.Conn)

	for {
		data, err := reader.ReadString('\n')

		if err == nil {
			switch t, msg := Parser(data); t {
			case newPlayer:
				fmt.Println(msg)

			}
		}
	}
}

// ConfirmUsername sets a username and informs the server about it
func (c *Client) ConfirmUsername(username string) {
	c.Name = username
	c.Conn.Write(Compose(newPlayer, username))
}
