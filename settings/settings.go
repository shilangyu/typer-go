package settings

import (
	"io/ioutil"
	"path"
	"runtime"

	"github.com/shilangyu/typer-go/utils"
	"gopkg.in/yaml.v2"
)

type settings struct {
	Highlight Highlight
}

// I contains current settings properties from settings.yaml
var I settings

func init() {
	_, currFile, _, _ := runtime.Caller(0)
	settingsPath := path.Join(currFile, "..", "..", "settings.yaml")
	content, err := ioutil.ReadFile(settingsPath)
	utils.Check(err)
	err = yaml.Unmarshal(content, &I)
	utils.Check(err)
}
