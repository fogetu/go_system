package system_config

import (
	"github.com/astaxie/beego/config"
	"os"
	"path/filepath"
)

const MINE_POOL = "/mine/pool"

var (
	MinePoolAddr string
)

func init() {
	var appConfigPath string
	var filename string
	filename = "app.conf"
	if os.Getenv("BEEGO_RUNMODE") != "" {
		filename = os.Getenv("BEEGO_RUNMODE") + "." + filename
	}
	appConfigPath = filepath.Join("./", "conf", filename)
	var configErr error
	Iniconf, configErr = config.NewConfig("ini", appConfigPath)
	if configErr != nil {
		panic(configErr)
	}
	MinePoolAddr = Iniconf.String("MINE_POOL_ADDR")
}
