package system_config

import (
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	"os"
	"path/filepath"
	"strings"
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
	if lo := Iniconf.String("LogOutputs"); lo != "" {
		los := strings.Split(lo, ";")
		for _, v := range los {
			if logType2Config := strings.SplitN(v, ",", 2); len(logType2Config) == 2 {
				logs.SetLogger(logType2Config[0],logType2Config[1])
			} else {
				continue
			}
		}
	}
}

func Configer() config.Configer {
	return Iniconf
}
