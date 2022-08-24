package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SayHello() {
	r := gin.Default()
	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": fmt.Sprintf("Hello %s", name),
			"data":    []interface{}{},
		})
	})
	r.Run(":8888")
}
