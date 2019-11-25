package ui

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/shilangyu/typer-go/game"
	"github.com/shilangyu/typer-go/settings"
)

// CreateSingleplayer creates singleplayer screen widgets
func CreateSingleplayer(app *tview.Application) error {
	text, err := game.ChooseText()
	if err != nil {
		return err
	}
	state := game.NewState(text)

	statsWis := [...]*tview.TextView{
		tview.NewTextView().SetText("wpm: 0"),
		tview.NewTextView().SetText("time: 0s"),
	}

	pages := tview.NewPages().
		AddPage("modal", tview.NewModal().
			SetText("Play again?").
			AddButtons([]string{"yes", "exit"}).
			SetDoneFunc(func(index int, label string) {
				switch index {
				case 0:
					CreateSingleplayer(app)
				case 1:
					CreateWelcome(app)
				}
			}), false, false)

	var textWis []*tview.TextView
	for _, word := range state.Words {
		textWis = append(textWis, tview.NewTextView().SetText(word).SetDynamicColors(true))
	}

	currInput := ""
	inputWi := tview.NewInputField().
		SetFieldBackgroundColor(tcell.ColorDefault)
	inputWi.
		SetChangedFunc(func(text string) {
			if state.StartTime.IsZero() {
				state.Start()
				go func() {
					ticker := time.NewTicker(100 * time.Millisecond)
					for range ticker.C {
						if state.CurrWord == len(state.Words) {
							ticker.Stop()
							return
						}
						app.QueueUpdateDraw(func() {
							statsWis[0].SetText(fmt.Sprintf("wpm: %.0f", state.Wpm()))
							statsWis[1].SetText(fmt.Sprintf("time: %.02fs", time.Since(state.StartTime).Seconds()))
						})
					}
				}()
			}

			if len(currInput) < len(text) {
				if len(text) > len(state.Words[state.CurrWord]) || state.Words[state.CurrWord][len(text)-1] != text[len(text)-1] {
					state.IncError()
				}
			}

			app.QueueUpdateDraw(func(i int) func() {
				return func() {
					textWis[i].SetText(paintDiff(state.Words[i], text))
				}
			}(state.CurrWord))

			if text == state.Words[state.CurrWord] {
				state.NextWord()
				if state.CurrWord == len(state.Words) {
					state.End()
					app.QueueUpdateDraw(func() {
						pages.ShowPage("modal")

					})
				} else {
					inputWi.SetText("")
				}
			}

			currInput = text
		})

	flexTexts := tview.NewFlex()
	for _, textWi := range textWis {
		flexTexts.AddItem(textWi, len(textWi.GetText(true)), 1, false)
	}
	for _, statsWi := range statsWis {
		flexTexts.AddItem(statsWi, len(statsWi.GetText(true))+5, 1, false)
	}
	flexTexts.AddItem(inputWi, 0, 1, true)

	pages.AddPage("game", flexTexts, true, true).SendToBack("game")
	app.SetRoot(pages, true)

	keybindings(app, CreateWelcome)
	return nil
}

// takes a slice of words and length of a line
// returns xs and ys of the words on a plane
func organiseText(words []string, lineLength int) (points []struct{ x, y int }) {
	x, y := 0, 0

	for _, word := range words {
		if x+len(word) > lineLength {
			y++
			x = 0
		}
		points = append(points, struct{ x, y int }{x, y})
		x += len(word)
	}

	return
}

// paintDiff returns an tview-painted string displaying the difference
func paintDiff(toColor string, differ string) (colorText string) {
	var h string
	if settings.I.Highlight == settings.HighlightBackground {
		h = ":"
	}

	for i := range differ {
		if i >= len(toColor) || differ[i] != toColor[i] {
			colorText += "[" + h + "red]"
		} else {
			colorText += "[" + h + "green]"
		}

		switch settings.I.ErrorDisplay {
		case settings.ErrorDisplayTyped:
			colorText += string(differ[i])
		case settings.ErrorDisplayText:
			if i < len(toColor) {
				colorText += string(toColor[i])
			}
		}
	}
	colorText += "[-:-:-]"

	if len(differ) < len(toColor) {
		colorText += toColor[len(differ):]
	}

	return
}
