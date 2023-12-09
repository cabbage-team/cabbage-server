package controller

import (
	"cabbage-server/response"
	"cabbage-server/service"

	"github.com/gin-gonic/gin"
)

// PostSearch
// @Summary profile share
// @Description 个人主页分享
// @Tags post
// @Param name path string true "the user name"
// @Accept json
// @Router /v1/api/bio/{name} [get]
func ProfileSharre(ctx *gin.Context){
	name := ctx.Param("name")
	profile,err := service.ProfileShare(name)
	if err != nil {
		response.Error(ctx,err)
		return
	}
	response.Success(ctx,gin.H{
		"data":profile,
	})
}