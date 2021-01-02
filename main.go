/**
 * @Author: jiangbo
 * @Description:
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2021/01/02 1:26 下午
 */

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}

func main() {

	db := InitDB()
	defer db.Close()

	r := gin.Default()
	r.POST("/api/user/register", func(c *gin.Context) {
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
			name = RandomString(10)
		}
		log.Println(name, telephone, password)
		// 判断手机号是否存在
		if isTelephoneExist(db, telephone){
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": "422", "msg": "用户已经存在"})
			return
		}
		// 创建用户
		newUser := User{
			Name: name,
			Telephone: telephone,
			Password: password,
		}
		db.Create(&newUser)
		// 返回结果
		c.JSON(200, gin.H{
			"message": "注册成功",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func RandomString(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, n)
	// 设置随机种子，否则每次生产的随时数会是一样的
	rand.Seed(time.Now().Unix())
	for i, _ := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func isTelephoneExist(db *gorm.DB, telephone string) bool  {
	var user User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID !=0{
		return true
	}
	return false
}

func InitDB() *gorm.DB {

	driverName := "mysql"
	host := "192.168.8.109"
	port := "3306"
	database := "gin_jiangbo"
	username := "root"
	password := "123456"
	charset := "utf8"

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username, password, host, port, database, charset)

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("连接mysql失败，err:" + err.Error())
	}
	db.AutoMigrate(&User{})  // 自动创建数据表
	return db
}
