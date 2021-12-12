package router

import (
	"github.com/gin-gonic/gin"
	"jokerweb/controller/user"
)

func SetUp() *gin.Engine {

	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "ok")
	})
	r.Group("/user")
	{
		r.POST("login", user.UserLogin)
		r.POST("register", user.UserRegister)
	}
	return r
}
