package middleware

import (
	"bluebell/pkg/jwt"
	"bluebell/pkg/types"
	"github.com/gin-gonic/gin"
	"strings"
)

const CtxUserIDKey = "user_id"
const CtxUserNameKey = "user_name"

func JwtAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		//客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		//这里假设Token放在Header的Authorization中，并使用Bearer开头
		//Authorization: Bearer xxxx.xxx.xxx
		//具体实现根据具体业务决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			types.ResponseError(c, types.CodeInvalidToken)
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			types.ResponseError(c, types.CodeInvalidToken)
			c.Abort()
			return
		}
		claim, err := jwt.ParseToken(parts[1])
		if err != nil {
			types.ResponseError(c, types.CodeInvalidToken)
			c.Abort()
			return
		}
		c.Set(CtxUserIDKey, claim.UserID)
		c.Set(CtxUserNameKey, claim.Username)
		c.Next()
	}
}
