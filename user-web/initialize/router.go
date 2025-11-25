package initialize

import (
	"mxshop-api/user-web/middlewares"
	router2 "mxshop-api/user-web/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	//配置跨域
	Router.Use(middlewares.Cors())
	ApiGroup := Router.Group("/u/v1/")
	router2.InitUserRouter(ApiGroup)
	router2.InitBaseRouter(ApiGroup)
	//fmt.Print(Router)
	return Router
}
