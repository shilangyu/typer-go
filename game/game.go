package game

import (
	"errors"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/shilangyu/typer-go/settings"
)

// ChooseText randomly chooses a text from the dataset
func ChooseText() (string, error) {
	if _, err := os.Stat(settings.I.TextsPath); os.IsNotExist(err) {
		return "", errors.New("Didnt find typer texts, make sure your path is correct")
	}
	rand.Seed(time.Now().UTC().UnixNano())

	bytes, err := ioutil.ReadFile(settings.I.TextsPath)
	if err != nil {
		return "", errors.New("Couldnt load the typer texts, make sure the permission are correct")
	}
	texts := strings.Split(string(bytes), "\n")
	texts = texts[:len(texts)-1]

	return texts[rand.Intn(len(texts))], nil
}

// MessageType defines available commands
type MessageType string

const (
	// ChangeName is for changing usernames
	// payload = new username
	ChangeName MessageType = "change-name"
	// ExitPlayer is for players leaving the game
	// payload = nil
	ExitPlayer MessageType = "exit-player"
	// StartGame is for starting the game
	// payload = unix timestamp of when it starts
	StartGame MessageType = "start-game"
)

// Parse returns message type and message itself from a connection
func Parse(payload string) (MessageType, string) {
	s := strings.Split(payload[:len(payload)-1], ":")
	return MessageType(s[0]), strings.Join(s[1:], ":")
}

// Compose creates a data string for connections
func Compose(t MessageType, payload string) []byte {
	return []byte(string(t) + ":" + payload + "\n")
}
