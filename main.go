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

// @title BlueBell API Doc
// @version 1.0
// @description API Doc
// @termsOfService http://swagger.io/terms/

// @contact.name Jannan
// @contact.url https://www.jannan.top/
// @contact.email 1227891082@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:80
// @BasePath /

func main() {
	conf.Init()
	logger.Init()
	snowflake.Init(conf.Conf.StartTime, conf.Conf.MachineID)
	mysql.Init()
	redis.Init()
	validator.InitTrans("zh") // todo add config file

	logger.Log.Info("Swagger Doc in: http://127.0.0.1/swagger/index.html#/")

	r := router.Init(conf.Conf.Mode)
	_ = r.Run(fmt.Sprintf(":%d", conf.Conf.Port))
}
