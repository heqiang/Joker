package user

import "github.com/gin-gonic/gin"

func UserLogin(c *gin.Context) {
	c.String(200, "用户登录")
}

func UserRegister(c *gin.Context) {
	c.String(200, "用户注册")
}
