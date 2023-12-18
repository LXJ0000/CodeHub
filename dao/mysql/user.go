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
func (u *UserDao) CheckUserExist(username string) (bool, *models.UserModel) {
	var user *models.UserModel
	err := db.Where("user_name=?", username).First(&user).Error
	//err == nil 则用户名存在 则返回 true
	//否则返回false
	if err == nil {
		return true, user
	}
	return false, nil
}

// Create 添加用户
func (u *UserDao) Create(user *models.UserModel) error {
	return db.Create(&user).Error
}
