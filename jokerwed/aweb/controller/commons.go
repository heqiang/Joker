package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"jokerweb/utils"
)

func HandleValidtorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		utils.ResponseError(c, utils.CodeInvaildParam)
		return
	}
	//请求失败
	zap.L().Error("invaild param", zap.Error(err))
	utils.ResponseErrorWithMsg(c, RemoveTopStruct(errs.Translate(Trans)), utils.CodeInvaildParam)
	return
}
