package service

import (
	"bluebell/dao/mysql"
	"bluebell/pkg/logger"
	"bluebell/pkg/types"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
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
func (CommunityService) Info(c *gin.Context, rId string) {
	id, err := strconv.ParseInt(rId, 10, 64)
	if err != nil {
		types.ResponseError(c, types.CodeInvalidParams)
		logger.Log.Error("请求参数有误")
		return
	}
	var communityDao mysql.CommunityDao
	info, err := communityDao.GetInfo(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			types.ResponseError(c, types.CodeInvalidCommunityId)
			logger.Log.Error("查询失败，ID不存在")
		} else {
			types.ResponseError(c, types.CodeServerBusy)
			logger.Log.Error("服务繁忙")
		}
		return
	}
	types.ResponseSuccessWithData(c, info)
}
