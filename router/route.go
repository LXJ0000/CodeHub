package router

import (
	"bluebell/controllers"
	"bluebell/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	// 测试连接
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	// 404 接口
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	v1 := r.Group("/api/v1")
	{
		// 用户接口
		user := v1.Group("/user")
		{
			user.POST("/login", controllers.UserController{}.Login)
			user.POST("/register", controllers.UserController{}.Register)
		}

	}
	v1.Use(middleware.JwtAuthMiddleware())
	{
		// 社区接口
		community := v1.Group("/community")
		{
			community.GET("/", controllers.CommunityController{}.List)
		}

	}
	return r
}
