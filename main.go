package main

import (
	"context"
	"foglio/common"
	"foglio/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

var helpers common.Helpers

func main() {
	client := common.ConnectMongo()
	var err error
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			helpers.ErrorLogger.Fatal(err)
			return
		}
	}()

	app := gin.Default()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	app.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Foglio, the API!",
		})
	})

	router := app.Group("/api")
	routes.AddRouter(router)

	err = app.Run("localhost:3000")
	if err != nil {
		helpers.ErrorLogger.Fatal(err)
	}
}
