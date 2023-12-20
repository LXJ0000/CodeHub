package service

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/encrypt"
	"bluebell/pkg/jwt"
	"bluebell/pkg/logger"
	"bluebell/pkg/snowflake"
	"bluebell/pkg/types"
	"github.com/gin-gonic/gin"
)

type UserService struct {
}

func (u *UserService) Login(c *gin.Context, req *models.UserLoginReq) {
	userDao := mysql.NewUserDao()
	isExist, user := userDao.CheckUserExist(req.Username)
	if !isExist || !encrypt.CheckPassword(user.Password, req.Password) {
		logger.Log.Error("用户名或密码错误")
		types.ResponseError(c, types.CodeInvalidPassword)
		return
	}
	token, err := jwt.GenToken(user.UserID, user.UserName)
	if err != nil {
		logger.Log.Error("Token有误")
		types.ResponseError(c, types.CodeInvalidToken)
		return
	}
	types.ResponseSuccessWithToken(c, token)

}

func (u *UserService) Register(c *gin.Context, req *models.UserRegisterReq) {
	userDao := mysql.NewUserDao()
	//1. 判断用户是否存在
	if isExist, _ := userDao.CheckUserExist(req.Username); isExist {
		logger.Log.Error("用户名已存在")
		types.ResponseError(c, types.CodeUserExist)
		return
	}
	//2. 生成UID
	userID := snowflake.GenID()
	user := &models.UserModel{
		UserID:   userID,
		UserName: req.Username,
		Password: encrypt.GetPassword(req.Password),
	}
	//3. 入库
	if err := userDao.Create(user); err != nil {
		logger.Log.Error("用户注册失败")
		types.ResponseError(c, types.CodeServerBusy)
		return
	}
	types.ResponseSuccess(c)
}
