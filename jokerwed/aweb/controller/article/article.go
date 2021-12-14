package article

import (
	"github.com/gin-gonic/gin"
	"jokerweb/aweb/controller"
	"jokerweb/aweb/controller/article/articletype"
	"jokerweb/aweb/service/implements"
	"jokerweb/middlewares"
	"jokerweb/model"
	"jokerweb/utils"
)

func PostArticle(c *gin.Context) {
	var article model.Article
	if err := c.ShouldBind(&article); err != nil {
		controller.HandleValidtorError(c, err)
		return
	}
	var err error
	article.UserId, err = middlewares.GetCurrentUser(c)
	if err != nil {
		utils.ResponseError(c, utils.CodeNeedAuth)
		return
	}
	var art implements.Article
	err = art.PostArticle(article)
	if err != nil {
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(c, utils.CodeSuccess)
}

// UpdateArticle 更新mysql的同时更新es
func UpdateArticle(c *gin.Context) {
	var article articletype.Article
	if err := c.ShouldBind(&article); err != nil {
		controller.HandleValidtorError(c, err)
		return
	}
	var art implements.Article
	err := art.UpdateArticle(article)
	if err != nil {
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(c, utils.CodeSuccess)
}

func GetArticleById(c *gin.Context) {

}

func GetAllArticle(c *gin.Context) {

}
