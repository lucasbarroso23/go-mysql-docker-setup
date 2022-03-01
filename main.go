package main

import (
	"blog-service/infra"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // new gin router initialization
	router.GET("/", func(c *gin.Context) {
		infra.LoadEnv()     // loading env variables
		infra.NewDatabase() // new database connection
		c.JSON(http.StatusOK, gin.H{"data": "Hello blog"})
	})

	router.Run(":8000")
}
