package service

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
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
	if err := redis.CreatePostWithTime(post.PostID); err != nil {
		types.ResponseError(c, types.CodeServerBusy)
		logger.Log.Error("CreatePostWithTime ERROR")
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

func (PostService) ListPro(c *gin.Context, req *models.PostListProReq) {
	//1. 查询post_id
	postDao := redis.NewPostDao()
	ids, err := postDao.GetPostIDInorder(req)
	if err != nil {
		logger.Log.Error("ID 列表查询失败")
		types.ResponseError(c, types.CodeServerBusy)
		return
	}
	//2. 数据库查询详细信息
	postSqlDao := mysql.NewPostDao()
	list, err := postSqlDao.GetPostListWithIDList(ids)
	if err != nil {
		logger.Log.Error("数据库帖子列表查询失败")
		types.ResponseError(c, types.CodeServerBusy)
		return
	}

	var posts []*models.PostInfoResp
	authorDao := mysql.NewUserDao()
	communityDao := mysql.NewCommunityDao()
	for _, post := range list {
		authorName, _ := authorDao.GetUserName(post.AuthorID)
		community, _ := communityDao.GetInfo(post.CommunityID)
		info := &models.PostInfoResp{
			AuthorName:        authorName,
			PostResp:          post,
			CommunityInfoResp: community,
		}
		posts = append(posts, info)
	}

	//
	total, _ := postSqlDao.GetCountByCondition(nil)

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

// Vote for Post
func (PostService) Vote(c *gin.Context, userID int64, req *models.VoteReq) {
	redis.VoteForPost(c, userID, req)
}
