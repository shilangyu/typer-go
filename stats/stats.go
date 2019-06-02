package stats

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/shilangyu/typer-go/utils"
)

type history struct {
	Timestamp time.Time `json:"timestamp"`
	Wpm       float64   `json:"wpm"`
}

type stats struct {
	History []history `json:"history"`
}

// I contains current statistics
var I stats
var statsPath string

func init() {
	statsPath = path.Join(utils.Root(), "stats.json")
	content, err := ioutil.ReadFile(statsPath)
	utils.Check(err)
	err = json.Unmarshal(content, &I)
	utils.Check(err)
}

// Save saves the current statistics
func Save() error {
	bytes, err := json.Marshal(I)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(statsPath, bytes, os.ModePerm)
}

// AddHistory appends wpm with a timestamp to the history property
func AddHistory(wpm float64) {
	I.History = append(I.History, history{time.Now(), wpm})
}
