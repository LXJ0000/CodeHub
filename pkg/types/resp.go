package types

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code Code        `json:"code,omitempty"`
	Msg  interface{} `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type ListResp struct {
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}

func ResponseError(c *gin.Context, code Code) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  getMsg(code),
	})
}

func ResponseErrorWithMsg(c *gin.Context, code Code, msg interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  msg,
	})
}

func ResponseSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, &Response{
		Code: CodeSUCCESS,
		Msg:  getMsg(CodeSUCCESS),
	})
}
func ResponseSuccessWithToken(c *gin.Context, token string) {
	c.JSON(http.StatusOK, &Response{
		Code: CodeSUCCESS,
		Msg:  getMsg(CodeSUCCESS),
		Data: gin.H{
			"token": token,
		},
	})
}
func ResponseSuccessWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: CodeSUCCESS,
		Msg:  getMsg(CodeSUCCESS),
		Data: data,
	})
}

func ResponseSuccessWithList(c *gin.Context, total int64, list interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: CodeSUCCESS,
		Msg:  getMsg(CodeSUCCESS),
		Data: ListResp{
			Total: total,
			Data:  list,
		},
	})
}
