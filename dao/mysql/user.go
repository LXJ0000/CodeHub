package mysql

import (
	"bluebell/models"
	"sync"
)

type UserDao struct {
}

var (
	userDao  *UserDao
	userOnce sync.Once
)

func NewUserDao() *UserDao {
	userOnce.Do(func() {
		userDao = &UserDao{}
	})
	return userDao
}

// CheckUserExist 判断用户名是否存在
func (u *UserDao) CheckUserExist(username string) bool {
	err := db.Where("user_name=?", username).First(&models.UserModel{}).Error
	//err == nil 则用户名存在 则返回 true
	//法则返回false
	return err == nil
}

func (u *UserDao) Create(user *models.UserModel) error {
	return db.Create(&user).Error
}
