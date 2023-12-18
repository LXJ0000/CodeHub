package mysql

import (
	"bluebell/conf"
	"bluebell/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

var db *gorm.DB

func Init() {
	m := conf.Conf.MySQLConfig
	dsn := m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("mysql 连接失败%s", err.Error()))
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)

	MakeMigration()
}

func MakeMigration() {
	if err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.UserModel{}); err != nil {
		panic(err)
	}
}
