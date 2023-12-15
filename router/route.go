package router

import (
	"bluebell/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() *gin.Engine {
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

	return r
}
