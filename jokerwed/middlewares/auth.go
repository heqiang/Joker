package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"jokerweb/utils"
	"jokerweb/utils/jwt"
)

const (
	CTXUSERUSERNAMEKEY = "username"
	CTXUSERIDKEY       = "userid"
)

var ErrorUserNotLogin = errors.New("用户未登录")

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("token")
		if authHeader == "" {
			utils.ResponseError(c, utils.CodeNeedAuth)
			c.Abort()
			return
		}
		mc, err := jwt.ParseToken(authHeader)
		if err != nil {
			utils.ResponseError(c, utils.CodeInvaildAuth)
			c.Abort()
			return
		}
		// 将当前请求的UserId信息保存到请求的上下文c上
		c.Set(CTXUSERUSERNAMEKEY, mc.Username)
		c.Set(CTXUSERIDKEY, mc.UserId)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息

	}
}

func GetCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CTXUSERIDKEY)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userId, ok := uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return userId, nil
}
