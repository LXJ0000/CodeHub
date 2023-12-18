package main

import (
	"bluebell/conf"
	"bluebell/controllers"
	"bluebell/dao/mysql"
	"bluebell/pkg/snowflake"
	"bluebell/router"
	"fmt"
)

func main() {
	conf.Init()
	snowflake.Init(conf.Conf.StartTime, conf.Conf.MachineID)
	mysql.Init()

	controllers.InitTrans("zh") // todo add config file
	r := router.Init()
	_ = r.Run(fmt.Sprintf(":%d", conf.Conf.Port))
}
