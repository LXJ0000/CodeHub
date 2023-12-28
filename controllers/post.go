package controllers

import (
	"bluebell/models"
	"bluebell/pkg/logger"
	"bluebell/pkg/types"
	valid "bluebell/pkg/validator"
	"bluebell/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PostController struct {
}

func (PostController) Create(c *gin.Context) {
	var req *models.PostCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			types.ResponseErrorWithMsg(c, types.CodeInvalidParams, valid.RemoveTopStruct(errs.Translate(valid.Trans)))
			return
		}
		types.ResponseError(c, types.CodeInvalidParams)
		return
	}
	authorId, err := getCurrentUser(c)
	if err != nil {
		types.ResponseError(c, types.CodeInvalidToken)
		logger.Log.Error(err.Error())
		return
	}

	var serv service.PostService
	serv.Create(c, req, authorId)
}

// List 帖子列表
// @Summary 帖子列表接口
// @Description 可按社区分类，按时间或分数排序查询帖子列表
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query models.PostListProReq false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _PostListResp
// @Router /api/v1/post [get]
func (PostController) List(c *gin.Context) {
	// GET /api/v1/post?page=1&size=10&order=time&community_id=5
	req := &models.PostListProReq{
		Page:  1,
		Size:  10,
		Order: models.OrderTime, // magic string
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		types.ResponseError(c, types.CodeInvalidParams)
		logger.Log.Error("请求参数有误")
		return
	}

	var serv service.PostService
	serv.List(c, req)
}

// Info 帖子详细信息
func (PostController) Info(c *gin.Context) {
	rId := c.Param("id")

	var serv service.PostService
	serv.Info(c, rId)
}

func (PostController) Vote(c *gin.Context) {
	var req *models.VoteReq
	if err := c.ShouldBindJSON(&req); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok { // 类型断言
			types.ResponseErrorWithMsg(c, types.CodeInvalidParams, valid.RemoveTopStruct(errs.Translate(valid.Trans)))
			return
		}
		types.ResponseError(c, types.CodeInvalidParams)
		return
	}
	userID, err := getCurrentUser(c)
	if err != nil {
		types.ResponseError(c, types.CodeInvalidToken)
		logger.Log.Error(err.Error())
		return
	}

	var serv service.PostService
	serv.Vote(c, userID, req)
}
