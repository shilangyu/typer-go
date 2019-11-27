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

// Center returns a new primitive which shows the provided primitive in its
// center, given the provided primitive's size.
// credits: https://github.com/rivo/tview/blob/master/demos/presentation/center.go
func Center(width, height int, p tview.Primitive) tview.Primitive {
	return tview.NewFlex().
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(p, height, 1, true).
			AddItem(tview.NewBox(), 0, 1, false), width, 1, true).
		AddItem(tview.NewBox(), 0, 1, false)
}
