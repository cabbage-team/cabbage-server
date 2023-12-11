package controller

import (
	"cabbage-server/dto"
	"cabbage-server/internal"
	"cabbage-server/response"
	"cabbage-server/service"

	"github.com/gin-gonic/gin"
)

// CommentCreate
// @Summary create new social platform
// @Description 添加新的社交平台
// @Tags admin
// @Param request body dto.AddPlatformDTO true "platforms"
// @Accept json
// @Router /v1/api/admin/profile/platform/add [post]
func AddPlatform(c *gin.Context) {
	names := &dto.AddPlatformDTO{}
	err := c.BindJSON(names)
	if err != nil {
		response.Error(c, internal.RequestParamsNotValidError)
		return
	}
	err = service.AddPlatform(names.Name...)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, gin.H{})
}

// CommentCreate
// @Summary get all social platform
// @Description 查看所有社交平台
// @Tags admin
// @Accept json
// @Router /v1/api/admin/profile/platform/all [get]
func GetAllPlatform(c *gin.Context) {
	list, err := service.GetAllPlatform()
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, list)
}

// CommentCreate
// @Summary get stat data
// @Description 统计平台信息 今日新增贴子数 评论数 用户数  当月新增用户数 评论数 用户数
// @Tags admin
// @Param request query dto.StatMonthDataDTO true "the month number"
// @Accept json
// @Router /v1/api/admin/stat/count [get]
func Count(c *gin.Context) {
	month := &dto.StatMonthDataDTO{}
	err := c.BindQuery(month)
	if err != nil {
		response.Error(c, internal.RequestParamsNotValidError)
		return
	}
	countData, err := service.Count(month.Month)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, countData)
}
