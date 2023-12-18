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

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	user := r.Group("/user")
	{
		user.POST("/login", controllers.UserController{}.Login)
		user.POST("/register", controllers.UserController{}.Register)
	}
	user.Use(middleware.JwtAuthMiddleware())
	{

	}

	return r
}
