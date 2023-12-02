package router

import (
	"cabbage-server/controller"
	_ "cabbage-server/docs"
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

	UserAPI := v1Router.Group("user")
	UserAPI.POST("create", controller.CreateAccount)
	UserAPI.GET("profile", controller.GetUserProfile)

	// 参考 https://dev.to/tags
	TagsAPI := v1Router.Group("tag")
	TagsAPI.GET("list", controller.ReadTags)     // 查询所有tag
	TagsAPI.POST("hide", controller.HideTag)     // 当前用户hide一个tag或者取消hide
	TagsAPI.POST("follow", controller.FollowTag) // 当前用户follow一个tag或者取消follow

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
