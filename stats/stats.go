package stats

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"github.com/shilangyu/typer-go/utils"
)

// wordStat describes the errors concerning a word
type wordStat struct {
	Duration   time.Duration `json:"duration"`
	ErrorCount int           `json:"errorCount"`
}

type history struct {
	Timestamp time.Time `json:"timestamp"`
	Wpm       float64   `json:"wpm"`
}

type stats struct {
	History []history             `json:"history"`
	Words   map[string][]wordStat `json:"words"`
}

// I contains current statistics
// its initialized because json.Marshal sees the properties
// as null pointers not empty objects as it should
var I = stats{
	History: []history{},
	Words:   map[string][]wordStat{},
}
var statsPath string

func init() {
	userConfigDir, err := os.UserConfigDir()
	utils.Check(err)

	statsPath = path.Join(userConfigDir, "typer-go", "stats.json")
	if _, err := os.Stat(statsPath); os.IsNotExist(err) {
		err := os.MkdirAll(path.Dir(statsPath), 0644)
		utils.Check(err)

		file, err := os.Create(statsPath)
		file.Close()
		utils.Check(err)
		Save()
	}

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
	return ioutil.WriteFile(statsPath, bytes, 0644)
}

// AddHistory appends wpm with a timestamp to the history property
func AddHistory(wpm float64) {
	I.History = append(I.History, history{time.Now(), wpm})
}

// AddWord appends stats about a word
func AddWord(rawWord string, time time.Duration, errorCount int) {
	word := strings.TrimFunc(strings.ToLower(rawWord), func(ch rune) bool {
		return ch == ' ' || ch == '.' || ch == ':' || ch == '?' || ch == '!'
	})
	I.Words[word] = append(I.Words[word], wordStat{time, errorCount})
}
