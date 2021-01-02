/**
 * @Author: jiangbo
 * @Description:
 * @File:  UserController
 * @Version: 1.0.0
 * @Date: 2021/01/02 4:17 下午
 */

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"jiangbo.com/gin_jiang/common"
	"jiangbo.com/gin_jiang/model"
	"jiangbo.com/gin_jiang/utils"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	// 数据验证
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": "422", "msg": "手机号必须为11位"})
		return
	}
	// 判断手机号是否存在
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": "422", "msg": "密码不能少于6位"})
		return
	}
	// 名称如果不存在，则给一个默认的10位字符串
	if len(name) == 0 {
		name = utils.RandomString(10)
	}
	log.Println(name, telephone, password)
	// 判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": "422", "msg": "用户已经存在"})
		return
	}
	// 创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	DB.Create(&newUser)
	// 返回结果
	c.JSON(200, gin.H{
		"message": "注册成功",
	})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
