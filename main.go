package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	app.Run(":8080")
}
