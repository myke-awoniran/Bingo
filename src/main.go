package main

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello from micheal Go's server", "Gin")
	})
	router.Run(":3000")
}
