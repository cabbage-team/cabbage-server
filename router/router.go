package router

import (
	"cabbage-server/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Accept",
			"Origin",
			"Referer",
			"Content-Type",
			"Content-Length",
			"X-Requested-With",
			"Accept-Encoding",
			"Access-Control-Request-Headers",
			"Access-Control-Request-Method",
			"User-Agent",
			"Tenantid",
			"X-CSRF-Token",
			"Authorization",
		},
		AllowCredentials: true,
	}))
	/* v1 */
	v1Router := r.Group("/v1")
	v1Router.POST("/create/account", controller.CreateAccount)
	v1Router.GET("/get/userprofile", controller.GetUserProfile)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
