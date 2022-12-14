package main

import (
	"backend/exampleData"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, exampleData.GetData())
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
