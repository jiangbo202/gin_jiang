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
)

func CollectRouters(r *gin.Engine) *gin.Engine {

	r.POST("/api/user/register", controller.Register)
	return r
}

