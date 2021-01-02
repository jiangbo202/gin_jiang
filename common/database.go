/**
 * @Author: jiangbo
 * @Description:
 * @File:  database.
 * @Version: 1.0.0
 * @Date: 2021/01/02 4:22 下午
 */

package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"jiangbo.com/gin_jiang/model"
)

var DB *gorm.DB

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
	db.AutoMigrate(&model.User{}) // 自动创建数据表

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
