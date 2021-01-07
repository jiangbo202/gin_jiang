/**
 * @Author: jiangbo
 * @Description:
 * @File:  sys_base.go
 * @Version: 1.0.0
 * @Date: 2021/01/06 10:29 下午
 */

package routers

import (
	"github.com/gin-gonic/gin"
	"jiangbo.com/gin_jiang/controller"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("user")
	{
		BaseRouter.POST("register", controller.Register)
		BaseRouter.POST("login", controller.Login)
	}
	return BaseRouter
}
