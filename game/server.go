package game

import (
	"bufio"
	"net"
)

// Other contains all information needed about other players
type Other struct {
	// Conn holds the connection to the user
	Conn net.Conn
	// Name is the username
	Name string
	// Progress indicates how far is the user
	Progress float64
}

// Server is a state manager for the server
type Server struct {
	// State is an inheritance to store State logic
	State
	// Name is the username
	Name string
	// Server is an instance of net.Listener to control server logic
	Server net.Listener
	// Others has connection to clients
	Others []Other
}

// Accept listens for connections
func (s *Server) Accept() {
	for {
		conn, err := s.Server.Accept()

		if err == nil {
			newOther := Other{
				Conn: conn,
			}
			s.Others = append(s.Others, newOther)

			go s.Listen(&newOther)
		}
	}
}

// Listen listens to messages from client
func (s *Server) Listen(other *Other) {
	reader := bufio.NewReader(other.Conn)

	for {
		data, err := reader.ReadString('\n')

		if err == nil {
			switch t, msg := Parse(data); t {
			case changeName:
				other.Name = msg
			}
		}
	}
}

// Inform messages all clients about something
func (s *Server) Inform(msg []byte) {
	for _, client := range s.Others {
		client.Conn.Write(msg)
	}
}
