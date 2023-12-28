package router

import (
	"bluebell/controllers"
	_ "bluebell/docs" // 千万不要忘了导入把你上一步生成的docs
	"bluebell/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
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
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

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
		// 用户接口
		user := v1.Group("/user")
		{
			user.GET("/", controllers.UserController{}.Info)
		}

		// 社区接口
		community := v1.Group("/community")
		{
			community.GET("/", controllers.CommunityController{}.List)
			community.GET("/:id", controllers.CommunityController{}.Info)
		}

		//	帖子接口
		post := v1.Group("/post")
		{
			post.POST("/", controllers.PostController{}.Create)
			post.GET("/", controllers.PostController{}.List)
			post.GET("/:id", controllers.PostController{}.Info)
			post.POST("/vote", controllers.PostController{}.Vote)
		}
	}
	return r
}
