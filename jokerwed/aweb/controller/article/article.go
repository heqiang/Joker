package article

import (
	"github.com/gin-gonic/gin"
	"jokerweb/aweb/controller"
	"jokerweb/aweb/controller/article/articletype"
	"jokerweb/aweb/service/implements"
	"jokerweb/middlewares"
	"jokerweb/model"
	"jokerweb/utils"
	"strconv"
)

// PostArticle
// @Tags 文章
// @Summary 发布文章
// @title 发布文章
// @Security ApiKeyAuth
// @Param registerparam body  articletype.ArticlParam true "请求参数"
// @Success 200 {object} utils.ResponseData "{'code':200,'data':null,'msg':''}"
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
// @Success 200 {object} utils.ResponseData "{'code':200,'data':null,'msg':''}"
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
// @Success 200 {object} utils.ResponseData "{'code':200,'data':null,'msg':''}"
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
// @Param page query string false "页数"
// @Param size query string false "size"
// @Success 200 {object} utils.ResponseData "{'code':200,'data':null,'msg':''}"
// @Router /v1/user/getallarticle  [get]
func GetAllArticle(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}
	size, _ := strconv.Atoi(c.Query("size"))
	switch {
	case size > 100:
		size = 100
	case size <= 0:
		size = 10
	}
	offset := (page - 1) * size
	var art implements.Article
	allArticle, total, err := art.GetAllArticle(offset, size)
	if err != nil {
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	paginationQ := &model.ParamPostList{
		Total: total,
		Data:  allArticle,
	}
	utils.ResponseSuccessWithMsg(c, utils.CodeSuccess, paginationQ)
}

// VoteArticle
// @Tags 文章
// @Summary 文章点赞
// @title 文章点赞
// @Security ApiKeyAuth
// @Param vote body model.Vote true "请求参数"
// @Success 200 {object} utils.ResponseData "{'code':200,'data':null,'msg':''}"
// @Router /v1/user/voteArticle  [post]
func VoteArticle(c *gin.Context) {
	var vote model.Vote
	if err := c.ShouldBindJSON(&vote); err != nil {
		controller.HandleValidtorError(c, err)
		return
	}
	userId, err := middlewares.GetCurrentUser(c)
	if err != nil {
		utils.ResponseError(c, utils.CodeNeedAuth)
		return
	}
	var art implements.Article
	err = art.VoteArticle(vote, userId)
	if err != nil {
		utils.ResponseError(c, utils.CodeServerBusy)
	}
	utils.ResponseSuccess(c, utils.CodeSuccess)
}

func Comment(c *gin.Context) {
	var commentArticle articletype.CommentArticleParam
	var commentToComment articletype.CommentToCommentParam
	level, err := strconv.Atoi(c.Param("level"))
	if err != nil {
		utils.ResponseErrorWithMsg(c, "请输入必输参数level", utils.CodeInvaildParam)
		return
	}
	userid, _ := middlewares.GetCurrentUser(c)
	// 对文章进行评论
	if level == 1 {
		if err := c.ShouldBindJSON(&commentArticle); err != nil {
			controller.HandleValidtorError(c, err)
			return
		}
		var art implements.Article
		err := art.CommentArticle(commentArticle, userid)
		if err != nil {
			utils.ResponseError(c, utils.CodeServerBusy)
			return
		}
		utils.ResponseSuccess(c, utils.CodeSuccess)
		// 对评论进行评论
	} else if level == 2 {
		if err := c.ShouldBindJSON(&commentToComment); err != nil {
			controller.HandleValidtorError(c, err)
			return
		}
		var art implements.Article
		err := art.CommentToComment(commentToComment, userid)
		if err != nil {
			utils.ResponseError(c, utils.CodeServerBusy)
			return
		}
		utils.ResponseSuccess(c, utils.CodeSuccess)

	} else {
		utils.ResponseErrorWithMsg(c, "请输入必输参数level", utils.CodeInvaildParam)
		return
	}

}
