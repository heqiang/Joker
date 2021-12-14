package router

import (
	"github.com/gin-gonic/gin"
	user2 "jokerweb/aweb/controller/user"
)

func InitBaseRouter(r *gin.RouterGroup) {
	router := r.Group("base")
	{
		router.GET("/captcha", user2.GetCaptcha)
	}

}
