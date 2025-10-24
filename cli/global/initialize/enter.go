package initialize

import (
	"pusherGo/config"
	"pusherGo/global"
)

func GlobalInit() error {
	var err error
	global.Configs, err = config.LoadConfig("config_local.yaml")
	return err
}
