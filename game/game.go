package game

import (
	"errors"
	"io/ioutil"
	"math/rand"
	"os"
	"path"

	"strings"

	"github.com/shilangyu/typer-go/settings"
	"github.com/shilangyu/typer-go/utils"
)

// ChooseText randomly chooses a text from the dataset
func ChooseText() (string, error) {
	if _, err := os.Stat(settings.I.TextsPath); os.IsNotExist(err) {
		return "", errors.New("Didnt find typer texts, make sure your path is correct")
	}

	bytes, err := ioutil.ReadFile(path.Join(utils.Root(), "texts.txt"))
	if err != nil {
		return "", nil
	}
	content := string(bytes)
	texts := strings.Split(content, "\n")
	texts = texts[:len(texts)-1]

	return texts[rand.Intn(len(texts))], nil
}
