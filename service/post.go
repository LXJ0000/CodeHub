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

func (PostService) List(c *gin.Context, req *models.PostListReq) {
	postDao := mysql.NewPostDao()

	if req.Size == 0 {
		req.Size = 5
	}
	condition := make(map[string]interface{})
	if req.CommunityID != 0 {
		condition["community_id"] = req.CommunityID
	}
	total, err := postDao.GetCountByCondition(condition)
	if err != nil {
		logger.Log.Error("帖子数量查询失败")
		types.ResponseError(c, types.CodeServerBusy)
		return
	}

	list, err := postDao.GetList(condition, &req.Page)
	if err != nil {
		logger.Log.Error("帖子查询失败")
		types.ResponseError(c, types.CodeServerBusy)
		return
	}
	types.ResponseSuccessWithList(c, total, list)
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

func (PostService) Vote(c *gin.Context, userID int64, req *models.VoteReq) {
	/*
		基于用户投票的相关算法 https://ruanyifeng.com/blog/algorithm/
		投一票就加432分  86422/200 -> 200个赞则给帖子续一天 - 来自《Redis实战》

		- direction = 1
			- 没点过
			- 点过反对
		- direction = 0
			- 取消赞成
			- 取消反对
		- direction = -1
			- 没点过
			- 点过赞成

		投票的限制: 每个帖子自发表之日起一个星期内允许用户投票，超过一个星期不允许再投票
		1. 到期之后将Redis中保存的赞成票数以及反对票数存储到Mysql中
		2. 到期之后删除redis.KeyPostVotedPrefix
	*/

	types.ResponseSuccess(c)
}
