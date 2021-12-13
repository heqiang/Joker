package initlize

import (
	"github.com/gin-gonic/gin"
	"jokerweb/router"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	ApiGroup := r.Group("/v1")
	router.InitBaseRouter(ApiGroup)
	router.InitUserRouter(ApiGroup)
	return r

}
