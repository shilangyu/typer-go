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
