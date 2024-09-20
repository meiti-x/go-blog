package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/meiti-x/go-blog/internal/blog"
)

func MapBlogRoutes(group *gin.RouterGroup, h blog.Handlers) {
	group.POST("/create", h.Create())
	group.PUT("/:news_id", h.Update())
	group.DELETE("/:news_id", h.Delete())
	group.GET("/:news_id", h.GetByID())
	group.GET("/search", h.SearchByTitle())
	group.GET("", h.GetNews())
}
