package controllers

import (
	"bluebell/models"
	"bluebell/pkg/logger"
	"bluebell/pkg/types"
	valid "bluebell/pkg/validator"
	"bluebell/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
}

func (UserController) Login(c *gin.Context) {
	//1. 参数校验
	var req models.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Log.Error("请求参数有误")

		if errs, ok := err.(validator.ValidationErrors); ok {
			types.ResponseErrorWithMsg(c, types.CodeInvalidParams, valid.RemoveTopStruct(errs.Translate(valid.Trans)))
			return
		}
		types.ResponseError(c, types.CodeInvalidParams)
		return
	}
	//2. 业务处理
	var serv service.UserService
	serv.Login(c, &req)
}

func (UserController) Register(c *gin.Context) {
	//1. 参数获取校验
	var req models.UserRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Log.Error("请求参数有误")
		// 判断error是不是validator类型
		if errs, ok := err.(validator.ValidationErrors); ok {
			types.ResponseErrorWithMsg(c, types.CodeInvalidParams, valid.RemoveTopStruct(errs.Translate(valid.Trans)))
			return
		}
		types.ResponseError(c, types.CodeInvalidParams)
		return
	}
	//2. 业务处理
	var serv service.UserService
	serv.Register(c, &req)

}
