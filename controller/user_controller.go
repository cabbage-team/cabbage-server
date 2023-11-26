package controller

import (
	"cabbage-server/model"
	"cabbage-server/response"
	"cabbage-server/service"
	"github.com/gin-gonic/gin"
)

// CreateAccount 创建新用户
func CreateAccount(c *gin.Context) {
	var requestUser = model.User{}
	_ = c.Bind(&requestUser)
	service.CreateAccount(&requestUser)
}

// GetUserProfile 获取用户信息
func GetUserProfile(c *gin.Context) {
	email := c.Param("email")
	user, err := service.GetUserProfile(email)
	if err != nil {
		response.Fail(c, "请求失败", gin.H{"error": err.Error()})
		return
	}
	response.Success(c, gin.H{"data": user}, "请求成功")
}
