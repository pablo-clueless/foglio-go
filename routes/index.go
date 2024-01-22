package routes

import "github.com/gin-gonic/gin"

func AddRouter(router *gin.RouterGroup) {
	AuthRouter(router)
	UserRouter(router)
	JobRouter(router)
}
