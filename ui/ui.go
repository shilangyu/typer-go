package ui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func keybindings(app *tview.Application, goBack func(app *tview.Application) error) {
	if goBack != nil {
		app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyEsc {
				app.QueueUpdateDraw(func() {
					goBack(app)
				})
			}

			return event
		})
	}
}
