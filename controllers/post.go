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
	var serv service.PostService
	authorId, err := getCurrentUser(c)
	if err != nil {
		types.ResponseError(c, types.CodeInvalidToken)
		logger.Log.Error(err.Error())
		return
	}
	serv.Create(c, req, authorId)
}

func (PostController) List(c *gin.Context) {
	var req *models.PostListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		types.ResponseError(c, types.CodeInvalidParams)
		logger.Log.Info(*req)
		return
	}
	var serv service.PostService
	serv.List(c, req)
}

func (PostController) Info(c *gin.Context) {
	rId := c.Param("id")
	var serv service.PostService
	serv.Info(c, rId)
}

func (PostController) Vote(c *gin.Context) {
	//	参数校验
	var req *models.VoteReq
	if err := c.ShouldBindJSON(&req); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok { // 类型断言
			types.ResponseErrorWithMsg(c, types.CodeInvalidParams, valid.RemoveTopStruct(errs.Translate(valid.Trans)))
			return
		}
		types.ResponseError(c, types.CodeInvalidParams)
		logger.Log.Info(*req)
		return
	}
	//
	var serv service.PostService
	serv.Vote(c, req)
	//
}
