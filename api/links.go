package api

import "github.com/gin-gonic/gin"

type links struct {
	ID   int    `json: "id"`
	Name string `json: "name"`
	URL  string `json: "url"`
}

func listLinks(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": []links{
			{ID: 1, Name: "test", URL: "http://test.com"},
		},
	})
}
