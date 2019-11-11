package settings

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/shilangyu/typer-go/utils"
	"gopkg.in/yaml.v2"
)

type settings struct {
	Highlight    Highlight
	ErrorDisplay ErrorDisplay
	TextsPath    string
}

// I contains current settings properties from settings.yaml
var I settings
var settingsPath string

func init() {
	userConfigDir, err := os.UserConfigDir()
	utils.Check(err)

	settingsPath = path.Join(userConfigDir, "typer-go", "settings.yaml")
	if _, err := os.Stat(settingsPath); os.IsNotExist(err) {
		err := os.MkdirAll(path.Dir(settingsPath), os.ModePerm)
		utils.Check(err)

		file, err := os.Create(settingsPath)
		file.Close()
		utils.Check(err)
		Save()
	}

	content, err := ioutil.ReadFile(settingsPath)
	utils.Check(err)

	err = yaml.Unmarshal(content, &I)
	utils.Check(err)
}

// Save saves the current settings
func Save() error {
	bytes, err := yaml.Marshal(I)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(settingsPath, bytes, os.ModePerm)
}
