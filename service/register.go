/**
 * @Author: jiangbo
 * @Description:
 * @File:  register.go
 * @Version: 1.0.0
 * @Date: 2021/01/06 11:13 下午
 */

package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"jiangbo.com/gin_jiang/common"
	"jiangbo.com/gin_jiang/model/db_model"
	"jiangbo.com/gin_jiang/model/request"
	"log"
)

func Register(register *request.Register) (string, error) {

	db := common.GetDB()
	var user db_model.User
	db.Where("telephone=?", register.Telephone).First(&user)
	if user.ID != 0 {
		return "", errors.New("手机号已存在！")
	}

	// 创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	newUser := db_model.User{
		Name:      register.Name,
		Telephone: register.Telephone,
		Password:  string(hasedPassword),
	}
	DB := common.GetDB()
	DB.Create(&newUser)
	// 发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		log.Printf("token generate error : %v", err)
		return "", err
	}
	return token, nil
}

func Login(login *request.Login) (string, error) {

	db := common.GetDB()
	// 判断手机号是否存在
	var user db_model.User
	db.Where("telephone=?", login.Telephone).First(&user)
	if user.ID == 0 {
		return "", errors.New("用户不存在")
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		return "", errors.New("密码错误")
	}
	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		log.Printf("token generate error : %v", err)
		return "", errors.New("token生成错误")
	}
	return token, nil
}