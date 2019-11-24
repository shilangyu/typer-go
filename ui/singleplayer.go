package ui

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jroimartin/gocui"
	widgets "github.com/shilangyu/gocui-widgets"
	"github.com/shilangyu/typer-go/game"
)

// CreateSingleplayer creates singleplayer screen widgets
func CreateSingleplayer(g *gocui.Gui) error {
	text, err := game.ChooseText()
	if err != nil {
		return err
	}
	state := game.NewState(text)

	w, h := g.Size()

	statsFrameWi := widgets.NewCollection("singleplayer-stats", "STATS", false, 0, 0, w/5, h)

	statWis := []*widgets.Text{
		widgets.NewText("singleplayer-stats-wpm", "wpm: 0  ", false, false, 2, 1),
		widgets.NewText("singleplayer-stats-time", "time: 0s  ", false, false, 2, 2),
	}

	textFrameWi := widgets.NewCollection("singleplayer-text", "", false, w/5+1, 0, 4*w/5, 5*h/6+1)

	points := organiseText(state.Words, 4*w/5-2)
	var textWis []*widgets.Text
	for i, p := range points {
		textWis = append(textWis, widgets.NewText("singleplayer-text-"+strconv.Itoa(i), state.Words[i], false, false, w/5+1+p.x, p.y))
	}

	var inputWi *widgets.Input
	inputWi = widgets.NewInput("singleplayer-input", true, false, w/5+1, h-h/6, w-w/5-1, h/6, func(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) bool {
		if key == gocui.KeyEnter || len(v.Buffer()) == 0 && ch == 0 {
			return false
		}

		if state.StartTime.IsZero() {
			state.Start()
			go func() {
				ticker := time.NewTicker(100 * time.Millisecond)
				for range ticker.C {
					if state.CurrWord == len(state.Words) {
						ticker.Stop()
						return
					}

					g.Update(func(g *gocui.Gui) error {
						err := statWis[1].ChangeText(
							fmt.Sprintf("time: %.02fs", time.Since(state.StartTime).Seconds()),
						)(g)
						if err != nil {
							return err
						}

						err = statWis[0].ChangeText(
							fmt.Sprintf("wpm: %.0f", state.Wpm()),
						)(g)
						if err != nil {
							return err
						}

						return nil
					})
				}
			}()
		}

		gocui.DefaultEditor.Edit(v, key, ch, mod)

		b := v.Buffer()[:len(v.Buffer())-1]

		if ch != 0 && (len(b) > len(state.Words[state.CurrWord]) || rune(state.Words[state.CurrWord][len(b)-1]) != ch) {
			state.IncError()
		}

		ansiWord := state.PaintDiff(b)

		g.Update(textWis[state.CurrWord].ChangeText(ansiWord))

		if b == state.Words[state.CurrWord] {
			state.NextWord()
			if state.CurrWord == len(state.Words) {
				state.End()

				var popupWi *widgets.Modal
				popupWi = widgets.NewModal("singleplayer-popup", "The end of the end\n is the end of times who craes", []string{"play", "quit"}, true, w/2, h/2, func(i int) {
					popupWi.Layout(g)
				}, func(i int) {
					switch i {
					case 0:
						CreateSingleplayer(g)
					case 1:
						// CreateWelcome(g)
					}
				})
				g.Update(func(g *gocui.Gui) error {
					popupWi.Layout(g)
					popupWi.Layout(g)
					g.SetCurrentView("singleplayer-popup")
					g.SetViewOnTop("singleplayer-popup")
					return nil
				})

			}
			g.Update(inputWi.ChangeText(""))
		}

		return false
	})

	var wis []gocui.Manager
	wis = append(wis, statsFrameWi)
	for _, stat := range statWis {
		wis = append(wis, stat)
	}
	wis = append(wis, textFrameWi)
	for _, text := range textWis {
		wis = append(wis, text)
	}
	wis = append(wis, inputWi)

	g.SetManager(wis...)

	g.Update(func(*gocui.Gui) error {
		g.SetCurrentView("singleplayer-input")
		return nil
	})

	return keybindings(g, CreateSingleplayer) //CreateWelcome)
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
