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

	infoItems := utils.Center([]string{
		"How your text should be highlighted",
		"What part should be displayed when you type an error",
		"Path to your typer texts (seperated with a new line)",
	})
	infoWi := widgets.NewText("settings-menu-info", infoItems[0], true, true, w/2, 3*h/4)

	menuItems := utils.Center([]string{"highlight", "error display", "texts path"})
	menuWi := widgets.NewMenu("settings-menu", menuItems, true, w/4, h/2, func(i int) {
		g.Update(infoWi.ChangeText(infoItems[i]))
		currSetting(g)(i)
	}, nil)

	g.SetManager(menuWi, infoWi)
	defer currSetting(g)(0)

	if err := keybindings(g, CreateWelcome); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		focusedView++
		_, err := g.SetCurrentView("settings-" + []string{"menu", "side"}[focusedView%2])
		return err
	}); err != nil {
		return err
	}

	return nil
}

func currSetting(g *gocui.Gui) func(i int) {
	return func(i int) {
		w, h := g.Size()

		g.DeleteKeybindings("settings-side")
		g.DeleteView("settings-side")

		x, y := 3*w/4, h/2
		var sideWi gocui.Manager
		changes := func(g *gocui.Gui) {}

		switch i {
		case 0:
			menuItems := utils.Center([]string{
				settings.HighlightBackground.String(),
				settings.HighlightText.String(),
			})
			tempSideWi := widgets.NewMenu("settings-side", menuItems, true, x, y, func(i int) {
				settings.I.Highlight = settings.Highlight(i)
				settings.Save()
			}, nil)
			changes = func(g *gocui.Gui) { tempSideWi.ChangeSelected(int(settings.I.Highlight))(g) }

			sideWi = tempSideWi
		case 1:
			menuItems := utils.Center([]string{
				settings.ErrorDisplayTyped.String(),
				settings.ErrorDisplayText.String(),
			})
			tempSideWi := widgets.NewMenu("settings-side", menuItems, true, x, y, func(i int) {
				settings.I.ErrorDisplay = settings.ErrorDisplay(i)
				settings.Save()
			}, nil)
			changes = func(g *gocui.Gui) { tempSideWi.ChangeSelected(int(settings.I.ErrorDisplay))(g) }

			sideWi = tempSideWi
		}

		g.Update(func(g *gocui.Gui) error {
			g.SetCurrentView("settings-menu")
			sideWi.Layout(g)
			changes(g)
			return nil
		})
	}
}
