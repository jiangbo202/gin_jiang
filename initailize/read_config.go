/**
 * @Author: jiangbo
 * @Description:
 * @File:  read_config.go
 * @Version: 1.0.0
 * @Date: 2021/01/06 10:47 下午
 */

package initailize

import (
	"github.com/spf13/viper"
	"os"
)

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
