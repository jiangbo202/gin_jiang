/**
 * @Author: jiangbo
 * @Description:
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2021/01/02 1:26 下午
 */

package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"jiangbo.com/gin_jiang/common"
	"jiangbo.com/gin_jiang/initailize"
)

func main() {
	initailize.InitConfig()
	db := common.InitDB()
	defer db.Close()

	r := initailize.Routers()

	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}


