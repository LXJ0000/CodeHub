package main

import (
	"bluebell/conf"
	"bluebell/pkg/snowflake"
	"bluebell/router"
	"fmt"
)

func main() {
	conf.Init()
	snowflake.Init(conf.Conf.StartTime, conf.Conf.MachineID)

	r := router.Init()
	_ = r.Run(fmt.Sprintf(":%d", conf.Conf.Port))
}
