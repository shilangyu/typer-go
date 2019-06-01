package ui

import (
	"github.com/jroimartin/gocui"
	widgets "github.com/shilangyu/gocui-widgets"
	"github.com/shilangyu/typer-go/utils"
)

// CreateSettings creates a screen with settings
func CreateSettings(g *gocui.Gui) error {
	w, h := g.Size()
	g.Mouse = true

	infoItems := utils.Center([]string{
		"Single player mode - test your typing skills offline!",
		"Multi player mode - battle against other typers",
		"Settings - change app settings",
		"Exit - exit the app",
	})
	infoWi := widgets.NewText("welcome-menu-info", infoItems[0], true, true, w/2, 3*h/4)

	menuItems := []string{"highlight"}
	menuWi := widgets.NewMenu("welcome-main-menu", utils.Center(menuItems), true, true, w/2, h/2, func(i int) {
		g.Update(infoWi.ChangeText(infoItems[i]))
	}, nil)

	g.SetManager(menuWi, infoWi)

	if err := keybindings(g, CreateWelcome); err != nil {
		return err
	}

	return nil
}
