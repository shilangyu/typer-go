package game

import (
	"bufio"
	"fmt"
	"net"
	"time"
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
	// State stores State logic
	State *State
	// Name is the username
	Name string
	// Server is an instance of net.Listener to control server logic
	Server net.Listener
	// Others has connection to clients
	Others []*Other
	// Callback is fired when a message comes
	Callback func(t MessageType)
}

// Accept listens for connections
func (s *Server) Accept() {
	for {
		conn, err := s.Server.Accept()

		if err == nil {
			newOther := &Other{
				Conn: conn,
			}
			s.Others = append(s.Others, newOther)

			go s.Listen(newOther)
		}
	}
}

// Subscribe creates a callback
func (s *Server) Subscribe(cb func(t MessageType)) {
	s.Callback = cb
}

// Listen listens to messages from client
func (s *Server) Listen(other *Other) {
	reader := bufio.NewReader(other.Conn)

	for {
		data, err := reader.ReadString('\n')

		if err != nil {
			other.Conn.Close()
			for i, client := range s.Others {
				if client == other {
					s.Others = append(s.Others[:i], s.Others[i+1:]...)

					if s.Callback != nil {
						s.Callback(ExitPlayer)
					}

					break
				}
			}
			return
		}

		t, payload := Parse(data)

		switch t {
		case ChangeName:
			other.Name = payload
		}

		if s.Callback != nil {
			s.Callback(t)
		}

	}
}

// StartGame informs client the game is starting
func (s *Server) StartGame() {
	s.Inform(Compose(StartGame, fmt.Sprintf("%d", time.Now().Add(time.Second*5).Unix())))
}

// Inform messages all clients about something
func (s *Server) Inform(msg []byte) {
	for _, client := range s.Others {
		client.Conn.Write(msg)
	}
}
