package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Map news routes
func (s *Server) MapNewsRoutes(postsGroup *gin.Engine) error {
	v1 := postsGroup.Group("/api/v1")

	v1.GET("/posts", func(context *gin.Context) {
		fmt.Println("Getting posts...")
	})
	return nil
}
