package service

import "bluebell/models"

type UserService struct {
}

func (u *UserService) Login() {
}

func (u *UserService) Register(req *models.UserRegisterRequest) {
	//1. 判断用户是否存在
	//2. 生成UID
	//3. 入库
}
