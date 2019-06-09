package game

import (
	"strings"
	"time"

	"github.com/shilangyu/typer-go/settings"
)

// State describes the state of a game
type State struct {
	// CurrWord is an index to State.Words
	CurrWord int
	// Words contains the text split by spaces
	Words []string
	// StartTime is a timestamp of the first keystroke
	StartTime time.Time
}

// NewState initializes State
func NewState(text string) *State {
	words := strings.Split(text, " ")
	for i := range words[:len(words)-1] {
		words[i] += " "
	}

	return &State{
		Words: words,
	}
}

// Wpm is the words per minute
func (s State) Wpm() float64 {
	return float64(s.CurrWord) / time.Since(s.StartTime).Minutes()
}

// PaintDiff returns an ANSII-painted string displaying the difference
func (s *State) PaintDiff(differ string) (ansiWord string) {
	var h string
	switch settings.I.Highlight {
	case settings.HighlightBackground:
		h = "4"
	case settings.HighlightText:
		h = "3"
	}

	toColor := s.Words[s.CurrWord]
	for i := range differ {
		if i >= len(toColor) || differ[i] != toColor[i] {
			ansiWord += "\u001b[" + h + "1m"
		} else {
			ansiWord += "\u001b[" + h + "2m"
		}
		ansiWord += string(differ[i])
	}
	ansiWord += "\u001b[0m"

	if len(differ) < len(toColor) {
		ansiWord += toColor[len(differ):]
	}

	return
}
