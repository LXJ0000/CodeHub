package main

import (
	"bluebell/conf"
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/pkg/logger"
	"bluebell/pkg/snowflake"
	"bluebell/pkg/validator"
	"bluebell/router"
	"fmt"
)

func main() {
	conf.Init()
	logger.Init()
	snowflake.Init(conf.Conf.StartTime, conf.Conf.MachineID)
	mysql.Init()
	redis.Init()
	validator.InitTrans("zh") // todo add config file

	r := router.Init(conf.Conf.Mode)
	_ = r.Run(fmt.Sprintf(":%d", conf.Conf.Port))
}
