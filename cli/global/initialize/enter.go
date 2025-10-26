package initialize

import (
	"pusherGo/config"
	"pusherGo/global"
)

func GlobalInit() error {
	var err error
	global.Configs, err = config.LoadConfig("/home/ubuntu/pusher-go/cli/config/config.local.yaml")
	return err
}
