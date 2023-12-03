package middleware

import (
	"github.com/gin-gonic/gin"
	limit "github.com/gin-contrib/size"
)

func RequestLimit() gin.HandlerFunc{
	return limit.RequestSizeLimiter(20)
}