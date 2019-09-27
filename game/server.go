package game

import "net"

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
		}
	}
}
