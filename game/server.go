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

// Listen listens for connections
func (s *Server) Listen() {
	for {
		conn, err := s.Server.Accept()

		if err == nil {
			s.Others = append(s.Others, conn)

			go func() {
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
			}()
		}
	}
}
