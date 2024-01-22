package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.RouterGroup) *gin.RouterGroup {
	auth := router.Group("/auth")
	auth.POST("/login", Login)
	auth.POST("/verify", Verify)
	return auth
}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Login!"})
}

func Verify(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Verify!"})
}
