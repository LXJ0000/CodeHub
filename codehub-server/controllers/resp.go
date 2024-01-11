package controllers

import (
	"bluebell/models"
	"bluebell/pkg/types"
)

type _PostListResp struct {
	Code    types.Code             `json:"code"`    //业务响应状态码
	Message interface{}            `json:"message"` // 业务响应状态码
	Data    []*models.PostInfoResp `json:"data"`    //数据
}
