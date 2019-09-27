package game

import "net"

// Client is a state manager for the client
type Client struct {
	State
	Name string
	Conn net.Conn
}
