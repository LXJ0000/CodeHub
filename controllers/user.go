package controllers

import (
	"bluebell/models"
	"bluebell/pkg/logger"
	valid "bluebell/pkg/validator"
	"bluebell/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type UserController struct {
}

func (UserController) Login(c *gin.Context) {
	//1. 参数校验
	var req models.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Log.Error("请求参数有误")

		if errs, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": valid.RemoveTopStruct(errs.Translate(valid.Trans)), // 翻译错误
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	//2. 业务处理
	userService := service.UserService{}
	if err := userService.Login(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		logger.Log.Error(req.Username, ": 用户名或密码错误")
		return
	}
	//3.返回响应
	c.JSON(http.StatusOK, gin.H{"msg": "登陆成功"})
}

func (UserController) Register(c *gin.Context) {
	//1. 参数获取校验
	var req models.UserRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Log.Error("请求参数有误")
		// 判断error是不是validator类型
		if errs, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": valid.RemoveTopStruct(errs.Translate(valid.Trans)), // 翻译错误
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	//2. 业务处理
	userService := service.UserService{}
	if err := userService.Register(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		logger.Log.Error("用户注册失败")
		return
	}
	//3.返回响应
	c.JSON(http.StatusOK, gin.H{"msg": "注册成功"})
}
