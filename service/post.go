package service

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/logger"
	"bluebell/pkg/snowflake"
	"bluebell/pkg/types"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
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

func (PostService) List(c *gin.Context) {
	postDao := mysql.NewPostDao()
	list, err := postDao.GetList()
	if err != nil {
		logger.Log.Error("帖子查询失败")
		types.ResponseError(c, types.CodeServerBusy)
		return
	}
	types.ResponseSuccessWithData(c, list)
}
func (PostService) Info(c *gin.Context, rId string) {
	id, err := strconv.ParseInt(rId, 10, 64)
	if err != nil {
		logger.Log.Error("请求参数有误")
		types.ResponseError(c, types.CodeInvalidParams)
		return
	}

	postDao := mysql.NewPostDao()
	authorDao := mysql.NewUserDao()
	communityDao := mysql.NewCommunityDao()

	post, err := postDao.GetInfo(id)
	authorName, _ := authorDao.GetUserName(post.AuthorID)
	community, _ := communityDao.GetInfo(post.CommunityID)

	info := &models.PostInfoResp{
		AuthorName:        authorName,
		PostResp:          post,
		CommunityInfoResp: community,
	}

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Log.Error("查询失败，PostID不存在")
			types.ResponseError(c, types.CodeInvalidPostId)
		} else {
			types.ResponseError(c, types.CodeServerBusy)
			logger.Log.Error("服务繁忙")
		}
		return
	}
	types.ResponseSuccessWithData(c, info)
}
