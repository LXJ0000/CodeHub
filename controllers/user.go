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
	//if len(req.Username) == 0 || len(req.Password) == 0 || len(req.RePassword) == 0 || len(req.Password) != len(req.RePassword) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "请求参数有误",
	//	})
	//	return
	//}
	//2. 业务处理
	userService := service.UserService{}
	userService.Register(&req)
	//3.返回响应
	c.JSON(http.StatusOK, gin.H{"ok": "ok"})
}
