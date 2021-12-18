package user

import (
	"github.com/gin-gonic/gin"
	"jokerweb/aweb/controller"
	"jokerweb/aweb/service/implements"
	"jokerweb/model"
	"jokerweb/utils"
)

// UserLogin
// @Tags 用户
// @Summary 用户登录
// @title 用户登录
// @Param login body  swag.User true "请求参数"
// @Success 200 {object} utils.ResponseData "{'code':200,'data':null,'msg':''}"
// @Router /v1/user/login  [post]
func UserLogin(c *gin.Context) {
	var u model.User
	if err := c.ShouldBind(&u); err != nil {
		controller.HandleValidtorError(c, err)
		return
	}
	if !store.Verify(u.CaptchaId, u.Captcha, true) {
		utils.ResponseSuccess(c, utils.CodeCaptchaError)
		return
	}
	var user implements.User
	token, err := user.Login(&u)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1004,
			"msg":  err.Error(),
		})
		return
	}
	utils.ResponseSuccessWithMsg(c, utils.CodeSuccess, token)
}

// UserRegister
// @Tags 用户
// @Summary 用户注册
// @title 用户注册
// @Param register body  swag.User true "请求参数"
// @Success 200 {object} utils.ResponseData "{'code':200,'data':null,'msg':''}"
// @Router /v1/user/register  [post]
func UserRegister(c *gin.Context) {
	var p model.User
	if err := c.ShouldBind(&p); err != nil {
		controller.HandleValidtorError(c, err)
		return
	}
	if !store.Verify(p.CaptchaId, p.Captcha, true) {
		utils.ResponseSuccess(c, utils.CodeCaptchaError)
		return
	}
	var user implements.User
	err := user.Register(&p)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1005,
			"msg":  err.Error(),
		})
		return
	}
	utils.ResponseSuccess(c, utils.CodeSuccess)
}

func GetIndex(c *gin.Context) {
	c.String(200, "index")
}
