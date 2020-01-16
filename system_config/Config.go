package system_config

import (
	"github.com/astaxie/beego/config"
	"os"
	"path/filepath"
)

var (
	Iniconf config.Configer
)

func init() {
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var filename = "app.conf"
	if os.Getenv("BEEGO_RUNMODE") != "" {
		filename = os.Getenv("BEEGO_RUNMODE") + ".app.conf"
	}
	var appConfigPath string
	appConfigPath = filepath.Join(workPath, "conf", filename)
	var configErr error
	Iniconf, configErr = config.NewConfig("ini", appConfigPath)
	if configErr != nil {
		panic(configErr)
	}
}

func Configer() config.Configer {
	return Iniconf
}
