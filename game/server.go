package game

import (
	"github.com/kanopeld/go-socket"
)

// NewServer initializes a server that just broadcasts all events
func NewServer(port string) (*socket.Server, error) {
	s, err := socket.NewServer(":" + port)
	if err != nil {
		return nil, err
	}
	players := make(Players)

	s.On(socket.CONNECTION_NAME, func(c socket.Client) {
		c.On(ChangeName, func(data []byte) {
			ID, nickname := ExtractChangeName(string(data))
			players.Add(ID, nickname)
			c.Broadcast(ChangeName, data)

			for ID, p := range players {
				c.Emit(ChangeName, ID+":"+p.Nickname)
			}
		})

		onExit := func() {
			delete(players, c.ID())
			c.Broadcast(ExitPlayer, []byte(c.ID()))
		}
		c.On(socket.DISCONNECTION_NAME, onExit)
		c.On(ExitPlayer, onExit)
	})

	go s.Start()

	return s, nil
}
