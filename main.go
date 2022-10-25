package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "curl /dashboard/config to upload file",
		})
	})

	r.POST("/dashboard/config", uploadJSON)

	r.Run()

}
