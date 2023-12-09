package controller

import (
	"cabbage-server/boot"
	validate "cabbage-server/common/validate"
	"cabbage-server/dto"
	"cabbage-server/internal"
	"cabbage-server/response"
	"cabbage-server/service"

	"github.com/gin-gonic/gin"
)

// CreateAccount
// @Summary Create user
// @Description 创建新用户
// @Tags user
// @Accept json
// @Param request body dto.SignupDTO true "admin account"
// @Router /v1/api/user/create [post]
func CreateAccount(c *gin.Context) {
	account := &dto.SignupDTO{}
	err := c.BindJSON(account)
	if err != nil {
		response.Error(c,err)
	}
	errMsg := validate.Validators(account)
	if len(errMsg) != 0 {
		// process errMsg
		paramsError := internal.RequestParamsNotValidError
		paramsError.ErrorMsg = errMsg
		response.Error(c,paramsError)
		return
	}
	go boot.Emit.Emit("user:register",account.Name,account.Email)
	_,err = service.CreateAccount(account)
	if err != nil {
		response.Error(c,err)
		return
	}
	response.Success(c, nil, "注册成功")
}

// GetUserProfile 获取用户信息
// CreateAccount 
// @Summary get user profile
// @Description 获取用户信息
// @Tags user
// @Accept json
// @Param request query dto.UserProfileDTO true "admin account"
// @Router /v1/api/user/profile [get]
func GetUserProfile(c *gin.Context) {
	profile := dto.UserProfileDTO{}
	c.BindQuery(profile)
	errmsg := validate.Validators(&profile)
	if len(errmsg) != 0 {
		// 处理错误
	}
	user, err := service.GetUserProfile(profile.Email)
	if err != nil {
		response.Error(c,err)
	}
	response.Success(c, gin.H{"data": user}, "请求成功")
}

// Login 用户登录
// Login 
// @Summary user login
// @Description 用户登录
// @Tags user
// @Accept json
// @Param request body dto.LoginDTO true "the user account"
// @Router /v1/api/user/login [post]
func Login(c *gin.Context){

}