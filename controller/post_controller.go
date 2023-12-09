package controller

import (
	"cabbage-server/boot"

	"github.com/gin-gonic/gin"
)

// PostCreate
// @Summary create post
// @Description 发帖
// @Tags post
// @Accept json
// @Router /v1/api/post/create [post]
func PostCreate(c *gin.Context){
	go boot.Emit.Emit("post:create","创建帖子")
}

// PostSearch 
// @Summary search post
// @Description 搜索帖子、文章
// @Tags post
// @Accept json
// @Router /v1/api/post/search [get]
func PostSearch(c *gin.Context){
	go boot.Emit.Emit("post:search",nil)
}

// PostOperator 
// @Summary operator post
// @Description 操作帖子 赞 踩 收藏
// @Tags post
// @Accept json
// @Router /v1/api/post/operater [post]
func PostOperator(c *gin.Context){
	go boot.Emit.Emit("post:search","12","33")
}


// PostDelete
// @Summary operator post
// @Description 删除帖子
// @Tags post
// @Accept json
// @Router /v1/api/post/del [delete]
func PostDelete(c *gin.Context){
	go boot.Emit.Emit("post:delete", "","")
}
