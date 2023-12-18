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

func (u *UserDao) CheckUserExist(username string) bool {
	err := db.Model(&models.UserModel{}).Where("user_name=?", username).First(&models.UserModel{}).Error
	return err == nil
}

func (u *UserDao) Create(user *models.UserModel) error {
	return db.Create(&user).Error
}
