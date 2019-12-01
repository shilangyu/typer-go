package game

import (
	"strconv"
	"strings"
)

const (
	// ChangeName is for changing usernames
	// payload = ID:new username
	ChangeName = "change-name"
	// ExitPlayer is for players leaving the game
	// payload = ID
	ExitPlayer = "exit-player"
	// StartGame is for starting the game
	// payload = unix timestamp of when it starts
	StartGame = "start-game"
	// Progress is for indicating typing progress
	// payload = ID:%
	Progress = "progress"
)

// Events is an array of all available events
var Events = [...]string{ChangeName, ExitPlayer, StartGame, Progress}

func split(s string) (string, string) {
	ss := strings.SplitN(s, ":", 2)
	return ss[0], ss[1]
}

// ExtractChangeName takes a payload and gives extracted data
func ExtractChangeName(payload string) (ID, nickname string) {
	return split(payload)
}

// ExtractExitPlayer takes a payload and gives extracted data
func ExtractExitPlayer(payload string) (ID string) {
	return payload
}

// ExtractStartGame takes a payload and gives extracted data
func ExtractStartGame(payload string) (unixTimestamp int64) {
	unixTimestamp, _ = strconv.ParseInt(payload, 10, 64)
	return
}

// ExtractProgress takes a payload and gives extracted data
func ExtractProgress(payload string) (ID string, progress int) {
	ID, p := split(payload)
	progress, _ = strconv.Atoi(p)
	return
}
