package router

import (
	"github.com/gin-gonic/gin"
	"jokerweb/controller/user"
)

func InitBaseRouter(r *gin.RouterGroup) {
	router := r.Group("base")
	{
		router.GET("/captcha", user.GetCaptcha)
	}

}
