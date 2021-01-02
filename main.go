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
	"jiangbo.com/gin_jiang/common"
)


func main() {

	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = CollectRouters(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}





