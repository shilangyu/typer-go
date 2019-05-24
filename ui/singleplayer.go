package ui

import (
	"github.com/jroimartin/gocui"
	"github.com/shilangyu/typeracer-go/widgets"
)

// CreateSingleplayer creates welcome screen widgets
func CreateSingleplayer(g *gocui.Gui) error {
	w, h := g.Size()

	stats := [...]*widgets.Text{
		widgets.NewText("singleplayer-stats-wpm", "wpm: 0", false, true, w/10, h/10),
		widgets.NewText("singleplayer-stats-time", "time: 0", false, true, w/10, h/10+1),
	}

	text := widgets.NewText("singleplayer-text", "Cock and balls", true, false, w/8, h/8)

	g.SetManager(text, stats[0], stats[1])

	if err := keybindings(g); err != nil {
		return err
	}

	return nil
}
