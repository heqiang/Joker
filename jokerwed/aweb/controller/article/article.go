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

// PostArticle
// @Tags 文章
// @Summary 发布文章
// @title 发布文章
// @Security ApiKeyAuth
// @Param registerparam body  articletype.ArticlParam true "请求参数"
// @Success 200 {string} json '{ "code": 1000, "msg": "success"}'
// @Router /v1/user/postarticle  [post]
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

// UpdateArticle
// @Tags 文章
// @Summary 文章更新
// @title 文章更新
// @Security ApiKeyAuth
// @Param  updatearticle  body  articletype.ArticlParam true "文章更新参数"
// @Success 200 {string} json '{"code": 1000, "msg": "success"}'
// @Router /v1/user/updatearticle  [post]
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

// GetArticleById
// @Tags 文章
// @Summary 通过id获取文章
// @title 通过id获取文章
// @Security ApiKeyAuth
// @Param articleId path string true "articleId"
// @Success 200 {string} json '{ "code": 1000, "msg": "success"}'
// @Router /v1/user/getarticlebyid/{articleId}  [get]
func GetArticleById(c *gin.Context) {
	articleId := c.Param("articleId")
	if articleId == "" {
		utils.ResponseSuccessWithMsg(c, utils.CodeInvaildParam, "请输入正确的文章id")
		return
	}
	var art implements.Article
	article, err := art.GetArticleByArticleId(articleId)
	if err != nil {
		utils.ResponseError(c, utils.CodeArticleNotExist)
		return
	}
	utils.ResponseSuccessWithMsg(c, utils.CodeSuccess, article)
}

// GetAllArticle
// @Tags 文章
// @Summary 获取所有的文章
// @title 获取所有的文章
// @Security ApiKeyAuth
// @Success 200 {string} json '{"code": 1000, "msg": "success"}'
// @Router /v1/user/getallarticle  [get]
func GetAllArticle(c *gin.Context) {
	var art implements.Article
	allArticle := art.GetAllarticle()
	utils.ResponseSuccessWithMsg(c, utils.CodeSuccess, allArticle)
}
