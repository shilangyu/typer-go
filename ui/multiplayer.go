package ui

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jroimartin/gocui"
	widgets "github.com/shilangyu/gocui-widgets"
	"github.com/shilangyu/typer-go/game"
	"github.com/shilangyu/typer-go/utils"
)

const tcpPort = "9001"

var isHost bool
var server net.Listener
var conn net.Conn

// CreateMultiplayerSetup creates multiplayer room creation
func CreateMultiplayerSetup(g *gocui.Gui) error {
	w, h := g.Size()

	errorWi := widgets.NewText("mp-setup-error", strings.Repeat(" ", 59), false, true, w/2, h/4)

	infoItems := utils.Center([]string{"Be the host a type race - let your friends know your ip", "Join a room - enter the ip of the host"})
	infoWi := widgets.NewText("mp-setup-menu-info", infoItems[0], true, true, w/2, 3*h/4)

	menuItems := utils.Center([]string{"server", "client"})
	menuWi := widgets.NewMenu("mp-setup-menu", menuItems, true, w/4, h/2, func(i int) {
		g.Update(infoWi.ChangeText(infoItems[i]))
	}, func(i int) {
		switch i {
		case 0:
			isHost = true

			createWi := widgets.NewText("mp-setup-create", strings.Repeat(" ", w/4-3), true, true, 3*w/4+1, h/2)
			createWi.Layout(g)
			g.SetCurrentView("mp-setup-create")

			g.Update(createWi.ChangeText("Creating a room..."))

			conn, _ := net.Dial("udp", "8.8.8.8:80")
			localAddr := conn.LocalAddr().(*net.UDPAddr)
			myIP := localAddr.IP.String()
			conn.Close()

			tempServer, err := net.Listen("tcp", myIP+":"+tcpPort)

			if err != nil {
				g.Update(func(g *gocui.Gui) error {
					errorWi.ChangeText("\u001b[31mCould not create a server. Make sure the port 9001 is free.")(g)
					g.SetCurrentView("mp-setup-menu")
					return g.DeleteView("mp-setup-create")
				})
			} else {
				g.DeleteKeybindings("mp-setup-menu")

				server = tempServer
				g.Update(createWi.ChangeText(fmt.Sprintf("Room created at %s", myIP)))
				time.AfterFunc(2*time.Second, func() { CreateMultiplayer(g) })
			}

		case 1:
			isHost = false

			ipInputWi := widgets.NewInput("mp-setup-join", true, true, 3*w/4, h/2, w/4, 3, func(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) bool {
				if key == gocui.KeyEnter {
					IP := v.Buffer()[:len(v.Buffer())-1]
					tempConn, err := net.Dial("tcp", IP+":"+tcpPort)
					if err != nil {
						g.Update(func(g *gocui.Gui) error {
							errorWi.ChangeText("\u001b[31mCould not join the room. Please check your/servers firewall.")(g)
							g.SetCurrentView("mp-setup-menu")
							return g.DeleteView("mp-setup-join")
						})
					} else {
						conn = tempConn
						CreateMultiplayer(g)
					}
					return false
				}
				return !(len(v.Buffer()) == 0 && ch == 0)
			})
			ipInputWi.Layout(g)
			g.SetCurrentView("mp-setup-join")
		}
	})

	g.SetManager(infoWi, menuWi, errorWi)
	g.Update(func(*gocui.Gui) error {
		g.SetCurrentView("mp-setup-menu")
		menuWi.Layout(g)
		return nil
	})

	return keybindings(g, CreateWelcome)
}

// CreateMultiplayer creates multiplayer screen widgets
func CreateMultiplayer(g *gocui.Gui) error {
	text, err := game.ChooseText()
	if err != nil {
		return err
	}
	state := game.NewState(text)

	w, h := g.Size()

	statsFrameWi := widgets.NewCollection("multiplayer-stats", "STATS", false, 0, 0, w/5, h)

	statWis := []*widgets.Text{
		widgets.NewText("multiplayer-stats-wpm", "wpm: 0  ", false, false, 2, 1),
		widgets.NewText("multiplayer-stats-time", "time: 0s  ", false, false, 2, 2),
	}

	textFrameWi := widgets.NewCollection("multiplayer-text", "", false, w/5+1, 0, 4*w/5, 5*h/6+1)

	points := organiseText(state.Words, 4*w/5-2)
	var textWis []*widgets.Text
	for i, p := range points {
		textWis = append(textWis, widgets.NewText("multiplayer-text-"+strconv.Itoa(i), state.Words[i], false, false, w/5+1+p.x, p.y))
	}

	var inputWi *widgets.Input
	inputWi = widgets.NewInput("multiplayer-input", true, false, w/5+1, h-h/6, w-w/5-1, h/6, func(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) bool {
		if key == gocui.KeyEnter || len(v.Buffer()) == 0 && ch == 0 {
			return false
		}

		if state.StartTime.IsZero() {
			state.Start()
			go func() {
				ticker := time.NewTicker(100 * time.Millisecond)
				for range ticker.C {
					if state.CurrWord == len(state.Words) {
						ticker.Stop()
						return
					}

					g.Update(func(g *gocui.Gui) error {
						err := statWis[1].ChangeText(
							fmt.Sprintf("time: %.02fs", time.Since(state.StartTime).Seconds()),
						)(g)
						if err != nil {
							return err
						}

						err = statWis[0].ChangeText(
							fmt.Sprintf("wpm: %.0f", state.Wpm()),
						)(g)
						if err != nil {
							return err
						}

						return nil
					})
				}
			}()
		}

		gocui.DefaultEditor.Edit(v, key, ch, mod)

		b := v.Buffer()[:len(v.Buffer())-1]

		if ch != 0 && (len(b) > len(state.Words[state.CurrWord]) || rune(state.Words[state.CurrWord][len(b)-1]) != ch) {
			state.IncError()
		}

		ansiWord := state.PaintDiff(b)

		g.Update(textWis[state.CurrWord].ChangeText(ansiWord))

		if b == state.Words[state.CurrWord] {
			state.NextWord()
			if state.CurrWord == len(state.Words) {
				state.End()

				var popupWi *widgets.Modal
				popupWi = widgets.NewModal("multiplayer-popup", "The end of the end\n is the end of times who craes", []string{"play", "quit"}, true, w/2, h/2, func(i int) {
					popupWi.Layout(g)
				}, func(i int) {
					switch i {
					case 0:
						CreateSingleplayer(g)
					case 1:
						CreateWelcome(g)
					}
				})
				g.Update(func(g *gocui.Gui) error {
					popupWi.Layout(g)
					popupWi.Layout(g)
					g.SetCurrentView("multiplayer-popup")
					g.SetViewOnTop("multiplayer-popup")
					return nil
				})

			}
			g.Update(inputWi.ChangeText(""))
		}

		return false
	})

	var wis []gocui.Manager
	wis = append(wis, statsFrameWi)
	for _, stat := range statWis {
		wis = append(wis, stat)
	}
	wis = append(wis, textFrameWi)
	for _, text := range textWis {
		wis = append(wis, text)
	}
	wis = append(wis, inputWi)

	g.SetManager(wis...)

	g.Update(func(*gocui.Gui) error {
		g.SetCurrentView("multiplayer-input")
		return nil
	})

	return keybindings(g, CreateMultiplayerSetup)
}

func createServer() (*net.Listener, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip := localAddr.IP
	conn.Close()

	listener, err := net.Listen("tcp", ip.String()+":9001")
	if err != nil {
		return nil, err
	}

	for {
		// Listen for an incoming connection.
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("someone connected")
		// Handle connections in a new goroutine.
		go func(conn net.Conn) {
			reader := bufio.NewReader(conn)

			for {
				message, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("someone disconnected")
					conn.Close()
					return
				}

				fmt.Print("|" + strings.TrimSpace(message) + "|")
				// Send a response back to person contacting us.
				conn.Write([]byte("STOP\n"))
			}
		}(conn)
	}

}
