package router

import (
	"cabbage-server/controller"
	"cabbage-server/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.GET("/users/:email", controller.GetUserByEmail)

	return r
}
