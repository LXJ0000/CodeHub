package controllers

import (
	"bluebell/service"
	"github.com/gin-gonic/gin"
)

type CommunityController struct {
}

// List 查询社区列表（id、name）
func (CommunityController) List(c *gin.Context) {
	var serv service.CommunityService
	serv.List(c)
}

// Info 获取某个社区的详细信息
func (CommunityController) Info(c *gin.Context) {
	//	1. 获取社区id
	rId := c.Param("id")
	var serv service.CommunityService
	serv.Info(c, rId)
}
