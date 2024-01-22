package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JobRouter(router *gin.RouterGroup) *gin.RouterGroup {
	job := router.Group("/job")
	job.POST("/create", CreateJob)
	job.GET("/all", GetAllJobs)
	job.GET("/:id", GetJob)
	job.POST("/:id/apply", ApplyToJob)
	job.POST("/:id/update", UpdateJob)
	job.POST("/:id/delete", DeleteJob)
	return job
}

func CreateJob(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Create Job!"})
}

func GetAllJobs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get All Jobs!"})
}

func GetJob(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Get Job!" + id})
}

func ApplyToJob(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Apply to Job!" + id})
}

func UpdateJob(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Update Job!" + id})
}

func DeleteJob(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Delete Job!" + id})
}
