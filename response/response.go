package response

import (
	"cabbage-server/internal"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 响应信息
func Response(ctx *gin.Context, httpStatus int, code int, data any, msg string) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

// Success 成功响应信息
func Success(ctx *gin.Context, data any) {
	Response(ctx, http.StatusOK, 0, data, "OK")
}

func Fail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 400, data, msg)
}

// Error 错误响应信息
func Error(c *gin.Context, err error) {
	if serverErr, ok := err.(*internal.ServerError); ok {
		if len(serverErr.ErrorMsg) != 0 {
			c.JSON(serverErr.Status, gin.H{
				"code":  serverErr.Code,
				"message": serverErr.Msg,
				"error":   serverErr.ErrorMsg,
			})
		} else {
			c.JSON(serverErr.Status, gin.H{
				"code":  serverErr.Code,
				"message": serverErr.Msg,
			})
		}
	}else{
		Fail(c,gin.H{
		},"some thing wrong")
	}
}
