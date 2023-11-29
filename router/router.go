package router

import (
	"cabbage-server/controller"
	"cabbage-server/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	/* v1 */
	v1Router := r.Group("/v1")
	v1Router.POST("/create/account", controller.CreateAccount)
	v1Router.GET("/get/userprofile", controller.GetUserProfile)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
