package controllers

import (
	"bluebell/middleware"
	"errors"
	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("用户未登录")

func getCurrentUser(c *gin.Context) (err error, userID int64) {
	uid, ok := c.Get(middleware.CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
	}
	return
}
