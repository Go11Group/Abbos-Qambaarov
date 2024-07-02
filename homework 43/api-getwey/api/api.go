package routes

import (
	"mymod/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	hand := handlers.NewHandler()
	api := router.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/", hand.UserHandler)
		}
		user.GET("/:id", hand.UserHandler)
	}
}