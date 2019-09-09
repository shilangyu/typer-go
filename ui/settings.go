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

	infoItems := utils.Center([]string{
		"How your text should be highlighted",
		"What part should be displayed when you type an error",
	})
	infoWi := widgets.NewText("settings-menu-info", infoItems[0], true, true, w/2, 3*h/4)

	menuItems := utils.Center([]string{"highlight", "error display"})
	menuWi := widgets.NewMenu("settings-menu", menuItems, true, w/4, h/2, func(i int) {
		g.Update(infoWi.ChangeText(infoItems[i]))
		currSetting(g)(i)
	}, nil)

	g.SetManager(menuWi, infoWi)
	currSetting(g)(0)

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

		g.DeleteView("settings-sidemenu")
		g.DeleteKeybindings("settings-sidemenu")

		var menuItems []string
		var settingsChange func(i int)
		var selected int

		switch i {
		case 0:
			menuItems = []string{
				settings.HighlightBackground.String(),
				settings.HighlightText.String(),
			}
			settingsChange = func(i int) {
				settings.I.Highlight = settings.Highlight(i)
			}
			selected = int(settings.I.Highlight)
		case 1:
			menuItems = []string{
				settings.ErrorDisplayTyped.String(),
				settings.ErrorDisplayText.String(),
			}
			settingsChange = func(i int) {
				settings.I.ErrorDisplay = settings.ErrorDisplay(i)
			}
			selected = int(settings.I.ErrorDisplay)
		}

		sideMenuWi := widgets.NewMenu("settings-sidemenu", utils.Center(menuItems), true, 3*w/4, h/2, func(i int) {
			settingsChange(i)
			settings.Save()
		}, nil)
		g.Update(func(g *gocui.Gui) error {
			g.SetCurrentView("settings-menu")
			sideMenuWi.Layout(g)
			sideMenuWi.ChangeSelected(selected)(g)
			return nil
		})
	}
}
