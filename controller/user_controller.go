package controller

import (
	"cabbage-server/common/utils"
	validate "cabbage-server/common/validate"
	"cabbage-server/dto"
	"cabbage-server/internal"
	"cabbage-server/model"
	"cabbage-server/response"
	"cabbage-server/service"

	"github.com/gin-gonic/gin"
)

// CreateAccount 创建新用户
func CreateAccount(c *gin.Context) {
	defer response.Error(c)
	account := &dto.SignupDTO{}
	err := c.BindJSON(account)
	if err != nil {
		panic(err)
	}
	errMsg := validate.Validators(account)
	if len(errMsg) != 0 {
		// process errMsg
		paramsError := internal.RequestParamsNotValidError
		paramsError.ErrorMsg = errMsg
		panic(paramsError)
	}
	err = service.CreateAccount(account)
	if err != nil {
		panic(err)
	}
	response.Success(c, nil, "注册成功")
}

// GetUserProfile 获取用户信息
func GetUserProfile(c *gin.Context) {
	defer response.Error(c)
	email := c.Query("email")
	if !utils.VerifyEmail(email) {
		panic("email格式错误")
	}
	user, err := service.GetUserProfile(email)
	if err != nil {
		panic(err)
	}
	response.Success(c, gin.H{"data": user}, "请求成功")
}
