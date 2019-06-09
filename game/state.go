package game

import (
	"strings"
	"time"
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
