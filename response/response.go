package response

import (
	"cabbage-server/internal"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 响应信息
func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

// Success 成功响应信息
func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

func Fail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 400, data, msg)
}

// Error 错误响应信息
func Error(c *gin.Context) {
	err := recover()
	if serverErr, ok := err.(*internal.ServerError); ok {
		Response(c, serverErr.Status, serverErr.Code, nil, serverErr.Msg)
		return
	}
	Fail(c, nil, err.(string))
}
