package controllers

import "github.com/gin-gonic/gin"

type UserController struct {
}

func (UserController) Login(c *gin.Context) {
	//1. 参数校验
	//2. 业务处理
	//3.返回响应
}

func (UserController) Register(c *gin.Context) {

}
