package main

import (
	"bluebell/conf"
	"bluebell/pkg/snowflake"
)

func main() {
	conf.Init()
	snowflake.Init(conf.Conf.StartTime, conf.Conf.MachineID)

	//fmt.Println(snowflake.GenID())
}
