package router

import (
	"github.com/gin-gonic/gin"
	"jokerweb/aweb/controller/article"
	"jokerweb/aweb/controller/user"
	"jokerweb/middlewares"
)

func InitUserRouter(r *gin.RouterGroup) {

	router := r.Group("/user")
	router.Use(middlewares.Cors())
	{
		router.POST("login", user.UserLogin)
		router.POST("register", user.UserRegister)
		router.Use(middlewares.JWTAuthMiddleware())
		router.POST("postarticle", article.PostArticle)
		router.POST("updatearticle", article.UpdateArticle)
		router.GET("getarticlebyid/:articleId", article.GetArticleById)
		router.GET("getallarticle", article.GetAllArticle)
	}
}
