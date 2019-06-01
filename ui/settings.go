package ui

import (
	"github.com/jroimartin/gocui"
	widgets "github.com/shilangyu/gocui-widgets"
	"github.com/shilangyu/typer-go/settings"
	"github.com/shilangyu/typer-go/utils"
)

// CreateSettings creates a screen with settings
func CreateSettings(g *gocui.Gui) error {
	var focusedView int

	w, h := g.Size()
	g.Mouse = true
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	menuItems := []string{"highlight"}
	menuWi := widgets.NewMenu("settings-menu", utils.Center(menuItems), true, true, w/4, h/2, currSetting(g), nil)
	currSetting(g)(0)

	g.SetManager(menuWi)

	if err := keybindings(g, CreateWelcome); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		focusedView++
		_, err := g.SetCurrentView("settings-" + []string{"menu", "sidemenu"}[focusedView%2])
		return err
	}); err != nil {
		return err
	}

	return nil
}

func currSetting(g *gocui.Gui) func(i int) {
	return func(i int) {
		w, h := g.Size()

		switch i {
		case 0:
			menuItems := []string{
				settings.HighlightBackground.String(),
				settings.HighlightText.String(),
			}
			sideMenuWi := widgets.NewMenu("settings-sidemenu", utils.Center(menuItems), true, true, 3*w/4, h/2, func(i int) {
				settings.I.Highlight = settings.Highlight(i)
				settings.Save()
			}, nil)
			g.Update(func(g *gocui.Gui) error {
				sideMenuWi.Layout(g)
				sideMenuWi.ChangeSelected(int(settings.I.Highlight))(g)
				g.SetCurrentView("settings-menu")
				return nil
			})
		}
	}
}
