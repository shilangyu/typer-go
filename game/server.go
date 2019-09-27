package game

import (
	"bufio"
	"fmt"
	"net"
)

// Server is a state manager for the server
type Server struct {
	State
	Name   string
	Server net.Listener
	Others []net.Conn
}

// Accept listens for connections
func (s *Server) Accept() {
	for {
		conn, err := s.Server.Accept()

		if err == nil {
			s.Others = append(s.Others, conn)

			go s.Listen(conn)
		}
	}
}

// Listen listens to messages from client
func (s *Server) Listen(conn net.Conn) {
	reader := bufio.NewReader(conn)

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

// Inform messages all clients about something
func (s *Server) Inform(msg []byte) {
	for _, client := range s.Others {
		client.Write(msg)
	}
}
