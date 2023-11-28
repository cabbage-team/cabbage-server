package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
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
	})
}
