package controller

import "github.com/gin-gonic/gin"


// PostCreate 
// @Summary create post
// @Description 发帖
// @Tags post
// @Accept json
// @Router /v1/api/post/create [post]
func PostCreate(c *gin.Context){

}

// PostSearch 
// @Summary search post
// @Description 搜索帖子、文章
// @Tags post
// @Accept json
// @Router /v1/api/post/search [post]
func PostSearch(c *gin.Context){

}

// PostOperator 
// @Summary operator post
// @Description 操作帖子 赞 踩 收藏
// @Tags post
// @Accept json
// @Router /v1/api/post/operater [post]
func PostOperator(c *gin.Context){

}


// PostDelete
// @Summary operator post
// @Description 删除帖子
// @Tags post
// @Accept json
// @Router /v1/api/post/del [delete]
func PostDelete(c *gin.Context){

}
