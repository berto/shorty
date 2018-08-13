package api

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter creates a gin Engine Router
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", ping)
	v1 := router.Group("api/v1")
	{
		v1.GET("/links", listLinks)
	}

	return router
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
