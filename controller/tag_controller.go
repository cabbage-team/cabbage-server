package controller

import (
	"cabbage-server/dto"
	"cabbage-server/response"
	"cabbage-server/service"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// CreateAccount 
// @Summary get topic
// @Description 获取话题列表
// @Tags topic
// @Accept json
// @Router /v1/api/tag/list [get]
func ReadTags(c *gin.Context) {
	tags := service.ReadTags()
	response.Success(c, gin.H{"data": tags}, "请求成功")
}

// hide tag 
// @Summary hide tag
// @Description 隐藏话题标签
// @Tags topic
// @Accept json
// @Param request query dto.TagFollowDTO true "topic"
// @Router /v1/api/tag/hide [get]
func HideTag(c *gin.Context) {

}

// follow tag 
// @Summary follow tag
// @Description 关注话题
// @Tags topic
// @Accept json
// @Param request body dto.TagFollowDTO true "topic"
// @Router /v1/api/tag/create [post]
func FollowTag(c *gin.Context) {

}

// CreateAccount 
// @Summary create new topic
// @Description 创建新话题标签
// @Tags topic
// @Accept json
// @Param request body dto.TagDTO true "topic"
// @Router /v1/api/tag/create [post]
func CreateTag(c *gin.Context) {
	tag := dto.TagDTO{}
	err := c.BindJSON(tag)
	if err != nil {
		// 参数错误
	}
	validate := validator.New()
	err = validate.Struct(tag)
	if err != nil && errors.Is(err, validator.ValidationErrors{}) {
		// 处理参数异常错误
	}
	// 获取数据
	// service.CreateTag(tag)
	//
	response.Success(c, gin.H{}, "OK")
}
