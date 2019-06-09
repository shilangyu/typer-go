package game

import (
	"io/ioutil"
	"math/rand"
	"path"
	"strings"

	"github.com/shilangyu/typer-go/utils"
)

// ChooseText randomly chooses a text from the dataset
func ChooseText() (string, error) {
	bytes, err := ioutil.ReadFile(path.Join(utils.Root(), "texts.txt"))
	if err != nil {
		return "", nil
	}
	content := string(bytes)
	texts := strings.Split(content, "\n")
	texts = texts[:len(texts)-1]

	return texts[rand.Intn(len(texts))], nil
}
