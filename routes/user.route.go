package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) *gin.RouterGroup {
	user := router.Group("/user")
	user.GET("/all", GetAllUsers)
	user.GET("/:id", GetUser)
	user.POST("/:id/update", UpdateUser)
	user.POST("/:id/delete", DeleteUser)
	return user
}

func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get All Users!"})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Get User!" + id})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Update User!" + id})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Delete User!" + id})
}
