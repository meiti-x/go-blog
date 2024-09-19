package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Execute() {
	fmt.Println("blog")
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "blad",
		})
	})
	r.Run()
}
