package service

import (
	"bluebell/dao/mysql"
	"bluebell/pkg/logger"
	"bluebell/pkg/types"
	"github.com/gin-gonic/gin"
)

type CommunityService struct {
}

func (CommunityService) List(c *gin.Context) {
	var communityDao mysql.CommunityDao
	list, err := communityDao.GetList()
	if err != nil {
		logger.Log.Error("查询失败")
		types.ResponseError(c, types.CodeServerBusy)
		return
	}
	types.ResponseSuccessWithData(c, list)
}
