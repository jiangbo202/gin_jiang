/**
 * @Author: jiangbo
 * @Description:
 * @File:  util
 * @Version: 1.0.0
 * @Date: 2021/01/02 4:21 下午
 */

package utils

import (
	"math/rand"
	"time"
)

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
