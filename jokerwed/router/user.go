package router

import (
	"github.com/gin-gonic/gin"
	"jokerweb/controller/user"
	"jokerweb/middlewares"
)

func InitUserRouter(r *gin.RouterGroup) {

	router := r.Group("/user")
	router.Use(middlewares.Cors())
	{
		router.POST("login", user.UserLogin)
		router.POST("register", user.UserRegister)
		router.Use(middlewares.JWTAuthMiddleware())
		router.GET("/index", user.GetIndex)
	}
}
