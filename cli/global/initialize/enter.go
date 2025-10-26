package initialize

import (
	"fmt"
	"os"
	"pusherGo/config"
	"pusherGo/global"
)

func GlobalInit() error {
	dir, err := os.Getwd()
	if err != nil {
		panic("error getting pwd")
	}
	global.Configs, err = config.LoadConfig(fmt.Sprintf("%s/config/config.local.yaml", dir))
	return err
}
