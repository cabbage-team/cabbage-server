package controller

import (
	"cabbage-server/model"
	"cabbage-server/response"
	"cabbage-server/service"
	"cabbage-server/util"
	"github.com/gin-gonic/gin"
)

// CreateAccount 创建新用户
func CreateAccount(c *gin.Context) {
	defer response.Error(c)
	var requestUser = model.User{}
	err := c.Bind(&requestUser)
	if err != nil {
		panic(err)
	}
	name := requestUser.Name
	email := requestUser.Email
	if len(name) == 0 {
		name = util.RandomString(10)
		return
	}
	if !util.VerifyEmail(email) {
		panic("邮箱格式错误")
		return
	}
	err = service.CreateAccount(&requestUser)
	if err != nil {
		panic(err)
	}
	response.Success(c, nil, "注册成功")
}

// GetUserProfile 获取用户信息
func GetUserProfile(c *gin.Context) {
	defer response.Error(c)
	email := c.Query("email")
	if !util.VerifyEmail(email) {
		panic("email格式错误")
	}
	user, err := service.GetUserProfile(email)
	if err != nil {
		panic(err)
	}
	response.Success(c, gin.H{"data": user}, "请求成功")
}
