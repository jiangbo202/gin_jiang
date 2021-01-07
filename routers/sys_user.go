/**
 * @Author: jiangbo
 * @Description:
 * @File:  sys_user.go
 * @Version: 1.0.0
 * @Date: 2021/01/06 10:39 下午
 */

package routers

import (
	"github.com/gin-gonic/gin"
	"jiangbo.com/gin_jiang/controller"
)

func InitUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	userRouter := Router.Group("user")
	{
		userRouter.POST("info", controller.Info)
	}
	return userRouter
}
