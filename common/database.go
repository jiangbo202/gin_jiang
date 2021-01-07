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
	"github.com/spf13/viper"
	"jiangbo.com/gin_jiang/model/db_model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {

	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username, password, host, port, database, charset)

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("连接mysql失败，err:" + err.Error())
	}
	db.AutoMigrate(&db_model.User{}) // 自动创建数据表

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
