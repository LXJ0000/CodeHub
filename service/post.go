package service

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/logger"
	"bluebell/pkg/snowflake"
	"bluebell/pkg/types"
	"github.com/gin-gonic/gin"
)

type PostService struct {
}

func (PostService) Create(c *gin.Context, req *models.PostCreateReq, authorID int64) {
	post := &models.PostModel{
		Title:       req.Title,
		Content:     req.Content,
		PostID:      snowflake.GenID(),
		AuthorID:    authorID,
		CommunityID: req.CommunityID,
		Status:      1,
	}
	postDao := mysql.NewPostDao()
	if err := postDao.Create(post); err != nil {
		types.ResponseError(c, types.CodeServerBusy)
		logger.Log.Error("帖子创建失败")
		return
	}
	types.ResponseSuccess(c)
}
