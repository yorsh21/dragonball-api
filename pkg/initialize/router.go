package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(resources Resources) error {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/character", resources.CharacterHandler.Create)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Route not found"})
	})

	return router.Run(":8080")
}
