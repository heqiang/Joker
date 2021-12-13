package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"jokerweb/controller"
	"jokerweb/model"
	"jokerweb/service/implements"
	"jokerweb/utils"
)

func UserLogin(c *gin.Context) {
	var u model.User
	if err := c.ShouldBind(&u); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			utils.ResponseError(c, utils.CodeInvaildParam)
			return
		}
		zap.L().Error("登录失败", zap.Error(err))
		utils.ResponseErrorWithMsg(c, controller.RemoveTopStruct(errs.Translate(controller.Trans)), utils.CodeInvaildParam)
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

func UserRegister(c *gin.Context) {
	var p model.User
	if err := c.ShouldBind(&p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			utils.ResponseError(c, utils.CodeInvaildParam)
			return
		}
		//请求失败
		zap.L().Error("用户注册 invaild param", zap.Error(err))
		utils.ResponseErrorWithMsg(c, controller.RemoveTopStruct(errs.Translate(controller.Trans)), utils.CodeInvaildParam)
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
