package controller

import (
	"cabbage-server/boot"
	"cabbage-server/common/validate"
	"cabbage-server/dto"
	"cabbage-server/internal"
	"cabbage-server/response"
	"cabbage-server/service"

	"github.com/gin-gonic/gin"
)

// CommentCreate
// @Summary create comment
// @Description 创建评论
// @Tags comment
// @Param request body dto.CommentDTO true "comment content"
// @Accept json
// @Router /v1/api/comment/create [post]
func CommentCreate(c *gin.Context) {
	go boot.Emit.Emit("comment:create","")
	comment := &dto.CommentDTO{}
	c.BindJSON(comment)
	errMsg := validate.Validators(comment)
	if len(errMsg) != 0 {
		paramsError := internal.RequestParamsNotValidError
		paramsError.ErrorMsg = errMsg
		response.Error(c,paramsError)
		return
	}
	userid := 0
	err := service.CreatePostComment(int64(userid),comment)
	if err != nil {
		response.Error(c,err)
		return
	}
	response.Success(c,gin.H{})
}


// CommentReply 
// @Summary reply comment
// @Description 回复评论
// @Tags comment
// @Accept json
// @Router /v1/api/comment/reply [post]
func CommentReply(c *gin.Context) {
	go boot.Emit.Emit("comment:create","")
}

// CommentOperator
// @Summary comment operator
// @Description 操作评论(赞,踩,收藏,分享)
// @Tags comment
// @Accept json
// @Param request body dto.CommentOperatorDTO true "operator code"
// @Router /v1/api/comment/operater [post]
func CommentOperator(c *gin.Context){
	go boot.Emit.Emit("comment:operator","")
}

// CommentOperator
// @Summary comment delete
// @Description 删除评论
// @Tags comment
// @Accept json
// @Router /v1/api/comment/del [delete]
func CommentDelete(c *gin.Context){
	go boot.Emit.Emit("comment:delete","")
}

// CommentOperator
// @Summary get comment
// @Description 查看评论
// @Tags comment
// @Accept json
// @Router /v1/api/comment/view [get]
func CommentView(c *gin.Context){
	go boot.Emit.Emit("comment:view","","")
}