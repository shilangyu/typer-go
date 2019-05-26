package ui

import (
	"github.com/jroimartin/gocui"
	"github.com/shilangyu/gocui-widgets"
)

// CreateSingleplayer creates welcome screen widgets
func CreateSingleplayer(g *gocui.Gui) error {
	w, h := g.Size()

	statsFrame := widgets.NewCollection("singleplayer-stats", "STATS", false, 0, 0, w/5, h)

	stats := [...]*widgets.Text{
		widgets.NewText("singleplayer-stats-wpm", "wpm: 0", false, false, 2, 1),
		widgets.NewText("singleplayer-stats-time", "time: 0", false, false, 2, 2),
	}

	text := widgets.NewText("singleplayer-text", "Cock and balls", true, false, w/5+1, 0)

	input := widgets.NewInput("singleplayer-input", true, false, w/5+1, h-h/6, w-w/5-1, h/6)

	g.SetManager(text, input, statsFrame, stats[0], stats[1])

	if err := keybindings(g); err != nil {
		return err
	}

	return nil
}
