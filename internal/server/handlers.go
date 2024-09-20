package server

import (
	"github.com/gin-gonic/gin"
	"github.com/meiti-x/go-blog/internal/blog/delivery"
)

// MapRoutes for mapping all routes
func (s *Server) MapRoutes(g *gin.Engine) error {
	// route init
	v1 := g.Group("/api")

	// blog routes
	blogHandlers := delivery.NewBlogsHandlers()
	blogGroup := v1.Group("/v1/blogs")
	delivery.MapBlogRoutes(blogGroup, blogHandlers)

	// auth routes
	return nil
}
