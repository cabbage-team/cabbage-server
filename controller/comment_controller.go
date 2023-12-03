package controller

import (
	"github.com/gin-gonic/gin"
)

// CommentCreate 
// @Summary create comment
// @Description 创建评论
// @Tags comment
// @Accept json
// @Router /v1/api/comment/create [post]
func CommentCreate(c *gin.Context) {

}


// CommentReply 
// @Summary reply comment
// @Description 回复评论
// @Tags comment
// @Accept json
// @Router /v1/api/comment/reply [post]
func CommentReply(c *gin.Context) {
	
}

// CommentOperator
// @Summary comment operator
// @Description 操作评论(赞,踩,收藏,分享)
// @Tags comment
// @Accept json
// @Param request body dto.CommentOperatorDTO true "operator code"
// @Router /v1/api/comment/operater [post]
func CommentOperator(c *gin.Context){
	const (
		CommentOperatorLike = 1<< iota
		CommentOperatorDiss
		CommentOperatorFavorite
		CommentOperatorShare
	)
}

// CommentOperator
// @Summary comment delete
// @Description 删除评论
// @Tags comment
// @Accept json
// @Router /v1/api/comment/del [delete]
func CommentDelete(c *gin.Context){

}