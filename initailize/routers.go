/**
 * @Author: jiangbo
 * @Description:
 * @File:  routers.go
 * @Version: 1.0.0
 * @Date: 2021/01/06 10:26 下午
 */

package initailize

import (
	"github.com/gin-gonic/gin"
	"jiangbo.com/gin_jiang/middleware"
	"jiangbo.com/gin_jiang/routers"
)

func Routers() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	publicGroup := router.Group("api/v1")  // 注册基础功能路由 不做鉴权
	{
		routers.InitBaseRouter(publicGroup)
	}
	PrivateGroup := router.Group("api/v1")
	PrivateGroup.Use(middleware.AuthMiddlewar())  // 需要认证
	{
		routers.InitUserRouter(PrivateGroup)
	}
	return router
}
