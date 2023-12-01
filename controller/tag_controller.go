package controller

import (
	"cabbage-server/dto"
	"cabbage-server/response"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


func ReadTags(c *gin.Context) {

}

func HideTag(c *gin.Context) {

}

func FollowTag(c *gin.Context) {

}

func CreateTag(c *gin.Context) {
	tag := dto.TagDTO{}
	err := c.BindJSON(tag)
	if err != nil {
		// 参数错误
	}
	validate := validator.New()
	err = validate.Struct(tag)
	if err != nil && errors.Is(err,validator.ValidationErrors{}) {
		// 处理参数异常错误
	}
	// 获取数据
	// service.CreateTag(tag)
	// 
	response.Success(c,gin.H{},"OK")
}

