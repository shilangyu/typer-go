package game

import (
	"strings"
	"time"

	"github.com/shilangyu/typer-go/stats"
)

// State describes the state of a game
type State struct {
	// CurrWord is an index to State.Words
	CurrWord int
	// Words contains the text split by spaces
	Words []string
	// StartTime is a timestamp of the first keystroke
	StartTime time.Time
	// properties concerning current word
	wordStart  time.Time
	wordErrors int
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

// Start starts the mechanism
func (s *State) Start() {
	s.StartTime = time.Now()
	s.wordStart = s.StartTime
}

// End ends the mechanism
func (s *State) End() {
	stats.AddHistory(s.Wpm())
	stats.Save()
}

// Wpm is the words per minute
func (s State) Wpm() float64 {
	return float64(s.CurrWord) / time.Since(s.StartTime).Minutes()
}

// Progress returns a float in the (0;1) range represending the progress made
func (s State) Progress() float64 {
	return float64(s.CurrWord) / float64(len(s.Words))
}

// IncError increments the error count
func (s *State) IncError() {
	s.wordErrors++
}

// NextWord saves stats of the current word and increments the counter
func (s *State) NextWord() {
	stats.AddWord(s.Words[s.CurrWord], time.Since(s.wordStart), s.wordErrors)
	s.CurrWord++

	s.wordStart = time.Now()
	s.wordErrors = 0
}
