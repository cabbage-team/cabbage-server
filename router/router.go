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
	v1Router := r.Group("/v1/api")
	v1Router.POST("/create/account", controller.CreateAccount)
	v1Router.GET("/get/userprofile", controller.GetUserProfile)

	// 参考 https://dev.to/tags
	v1Router.GET("/read/tags", controller.ReadTags)    // 查询所有tag
	v1Router.POST("/hide/tag", controller.HideTag)     // 当前用户hide一个tag或者取消hide
	v1Router.POST("/follow/tag", controller.FollowTag) // 当前用户follow一个tag或者取消follow

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
