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
		response.Error(c, err)
	}
	errMsg := validate.Validators(account)
	if len(errMsg) != 0 {
		// process errMsg
		paramsError := internal.RequestParamsNotValidError
		paramsError.ErrorMsg = errMsg
		response.Error(c, paramsError)
		return
	}
	go boot.Emit.Emit("user:register", account.Name, account.Email)
	_, err = service.CreateAccount(account)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, gin.H{})
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
	profile := &dto.UserProfileDTO{}
	c.BindQuery(profile)
	errmsg := validate.Validators(profile)
	if len(errmsg) != 0 {
		paramsError := internal.RequestParamsNotValidError
		paramsError.ErrorMsg = errmsg
		response.Error(c, paramsError)
		return
	}
	user, err := service.GetUserProfile(profile.Email)
	if err != nil {
		response.Error(c, err)
	}
	response.Success(c, gin.H{"data": user})
}

// Login 用户登录
// Login
// @Summary user login
// @Description 用户登录
// @Tags user
// @Accept json
// @Param request body dto.LoginDTO true "the user account"
// @Router /v1/api/user/login [post]
func Login(c *gin.Context) {

}

// Login 用户昵称检查
// Login
// @Summary check user name
// @Description 检查昵称
// @Tags user
// @Accept json
// @Param request query dto.NickNameDTO true "the user account"
// @Router /v1/api/user/name/check [get]
func CheckNickName(c *gin.Context) {
	name := &dto.NickNameDTO{}
	err := c.BindQuery(name)
	if err != nil {
		response.Error(c, internal.RequestParamsNotValidError)
		return
	}
	err = service.CheckNickName(name.Name)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, gin.H{})
}

// Login 获取用户关注列表
// Login
// @Summary check user name
// @Description 获取用户关注列表
// @Tags user
// @Param request query dto.FollowList true "the user account"
// @Accept json
// @Router /v1/api/user/follows [get]
func UserFollows(c *gin.Context) {
	params := &dto.FollowListDTO{}
	err := c.BindQuery(params)
	if err != nil {
		response.Error(c, internal.RequestParamsNotValidError)
		return
	}
	errMsg := validate.Validators(params)
	if len(errMsg) != 0 {
		paramsError := internal.RequestParamsNotValidError
		paramsError.ErrorMsg = errMsg
		response.Error(c, paramsError)
		return
	}
	userlist, err := service.UserFollowList(params.UID, params.Page, params.Size)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, userlist)
}

// Login 获取用户黑名单列表
// Login
// @Summary check user name
// @Description 获取用户黑名单列表
// @Tags user
// @Param request query dto.FollowList true "the user account"
// @Accept json
// @Router /v1/api/user/blacklist [get]
func UserBlackList(c *gin.Context) {
	params := &dto.FollowListDTO{}
	err := c.BindQuery(params)
	if err != nil {
		response.Error(c, internal.RequestParamsNotValidError)
		return
	}
	errMsg := validate.Validators(params)
	if len(errMsg) != 0 {
		paramsError := internal.RequestParamsNotValidError
		paramsError.ErrorMsg = errMsg
		response.Error(c, paramsError)
		return
	}
	userlist, err := service.UserBlackList(params.UID, params.Page, params.Size)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, userlist)
}
