/**
 * @Author: jiangbo
 * @Description:
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2021/01/02 1:26 下午
 */

package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"jiangbo.com/gin_jiang/common"
	"os"
)

func main() {
	InitConfig()
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = CollectRouters(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
