package controllers

import (
	"bluebell/models"
	"bluebell/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type UserController struct {
}

func (UserController) Login(c *gin.Context) {
	//1. 参数校验
	//2. 业务处理
	//3.返回响应
}

func (UserController) Register(c *gin.Context) {
	//1. 参数校验
	var req models.UserRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// todo 日志
		// 判断error是不是validator类型
		if errs, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": removeTopStruct(errs.Translate(trans)), // 翻译错误
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
		return
	}
	//3.返回响应
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
