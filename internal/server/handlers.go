package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/meiti-x/go-blog/config"
	"github.com/meiti-x/go-blog/internal/blog/delivery"
	"github.com/meiti-x/go-blog/internal/blog/repository"
	"github.com/meiti-x/go-blog/internal/blog/usecase"
)

// MapRoutes for mapping all routes
func (s *Server) MapRoutes(g *gin.Engine, db *sqlx.DB, cfg *config.Config) error {
	// route init
	v1 := g.Group("/api")

	// blog routes
	blogRepository := repository.NewBlogRepo(db)
	blogUseCase := usecase.NewBlogUseCase(cfg, blogRepository)
	blogHandlers := delivery.NewBlogsHandlers(db, blogUseCase)
	blogGroup := v1.Group("/v1/blogs")
	delivery.MapBlogRoutes(blogGroup, blogHandlers)

	// auth routes
	return nil
}
