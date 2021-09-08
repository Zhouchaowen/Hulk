package router

import (
	"Hulk/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	interfaceConfigRouter := router.Group("/admin")
	{
		controllers.InterfaceConfigRegister(interfaceConfigRouter)
	}
	return router
}
