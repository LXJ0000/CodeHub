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
