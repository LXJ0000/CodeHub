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
