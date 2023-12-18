package service

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/encrypt"
	"bluebell/pkg/snowflake"
	"errors"
	"fmt"
)

type UserService struct {
}

func (u *UserService) Login() {
}

func (u *UserService) Register(req *models.UserRegisterRequest) error {
	userDao := mysql.NewUserDao()
	//1. 判断用户是否存在
	if isExist := userDao.CheckUserExist(req.Username); isExist {
		return errors.New("用户名已存在")
	}
	//2. 生成UID
	userID := snowflake.GenID()
	user := &models.UserModel{
		UserID:   userID,
		UserName: req.Username,
		Password: encrypt.GetPassword(req.Password),
	}
	//3. 入库
	fmt.Println(user)
	err := userDao.Create(user)
	return err
}
