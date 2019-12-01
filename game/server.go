package game

import (
	"github.com/kanopeld/go-socket"
	"github.com/shilangyu/typer-go/utils"
)

// NewServer initializes a server that just broadcasts all events
func NewServer(port string) *socket.Server {
	s, err := socket.NewServer(":" + port)
	utils.Check(err)

	s.On(socket.CONNECTION_NAME, func(c socket.Client) {
		for _, eventName := range Events {
			func(evtName string) {
				c.On(evtName, func(data []byte) {
					c.Broadcast(evtName, data)
				})
			}(eventName)
		}

		c.On(socket.DISCONNECTION_NAME, func() {
			c.Broadcast(ExitPlayer, []byte(c.ID()))
		})
	})

	go s.Start()

	return s
}
