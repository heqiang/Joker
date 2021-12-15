package initlize

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "jokerweb/docs"
	"jokerweb/router"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	ApiGroup := r.Group("/v1")
	router.InitBaseRouter(ApiGroup)
	router.InitUserRouter(ApiGroup)
	return r

}
