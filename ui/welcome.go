package ui

import (
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/jroimartin/gocui"
	widgets "github.com/shilangyu/gocui-widgets"
	"github.com/shilangyu/typer-go/utils"
)

// CreateWelcome creates welcome screen widgets
func CreateWelcome(g *gocui.Gui) error {
	w, h := g.Size()
	g.Mouse = true
	g.Cursor = false
	g.Highlight = false

	signWi := widgets.NewText("welcome-sign", figure.NewFigure("typer-go", "", false).String(), false, true, w/2, h/5)

	infoItems := utils.Center([]string{
		"Single player mode - test your typing skills offline!",
		"Multi player mode - battle against other typers",
		"Settings - change app settings",
		"Exit - exit the app",
	})
	infoWi := widgets.NewText("welcome-menu-info", infoItems[0], true, true, w/2, 3*h/4)

	menuItems := utils.Center([]string{"single player", "multi player", "settings", "exit"})
	menuWi := widgets.NewMenu("welcome-main-menu", menuItems, true, true, w/2, h/2, func(i int) {
		g.Update(infoWi.ChangeText(infoItems[i]))
	}, func(i int) {
		switch i {
		case 0:
			CreateSingleplayer(g)
		case 2:
			CreateSettings(g)
		case 3:
			g.Close()
			os.Exit(0)
		default:

		}
	})

	g.SetManager(signWi, menuWi, infoWi)

	return keybindings(g, nil)
}
