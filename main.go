package main

import (
	"bluebell/conf"
	"bluebell/controllers"
	"bluebell/pkg/snowflake"
	"bluebell/router"
	"fmt"
)

func main() {
	conf.Init()
	snowflake.Init(conf.Conf.StartTime, conf.Conf.MachineID)

	controllers.InitTrans("zh") // todo add config file
	r := router.Init()
	_ = r.Run(fmt.Sprintf(":%d", conf.Conf.Port))
}
