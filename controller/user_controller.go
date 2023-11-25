package controller

import (
	"cabbage-server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := service.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
