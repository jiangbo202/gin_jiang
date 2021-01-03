/**
 * @Author: jiangbo
 * @Description:
 * @File:  routers
 * @Version: 1.0.0
 * @Date: 2021/01/02 4:26 下午
 */

package main

import (
	"github.com/gin-gonic/gin"
	"jiangbo.com/gin_jiang/controller"
	"jiangbo.com/gin_jiang/middleware"
)

func CollectRouters(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/user/register", controller.Register)
	r.POST("/api/user/login", controller.Login)
	r.GET("api/user/info", middleware.AuthMiddlewar(), controller.Info)
	return r
}

