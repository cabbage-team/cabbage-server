package controller

import (
	"cabbage-server/model"
	"cabbage-server/response"
	"cabbage-server/service"
	"cabbage-server/internal/serverr"
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
	defer func(ctx *gin.Context) {
		err := recover()
		if serverErr, ok := err.(*serverr.ServerError); ok {
			if len(serverErr.ErrorMsg) != 0 {
				ctx.JSON(serverErr.Status, gin.H{
					"code":  serverErr.Code,
					"msg": serverErr.Msg,
					"error":   serverErr.ErrorMsg,
				})
			} else {
				ctx.JSON(serverErr.Status, gin.H{
					"code":  serverErr.Code,
					"msg": serverErr.Msg,
				})
			}
		}
	}(c)
	email := c.Param("email")
	user, err := service.GetUserProfile(email)
	if err != nil {
		panic(err)
	}
	response.Success(c, gin.H{"data": user}, "请求成功")
}
