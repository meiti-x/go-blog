package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/meiti-x/go-blog/internal/blog"
)

type newsHandlers struct {
}

func NewBlogsHandlers() blog.Handlers {
	return &newsHandlers{}
}

func (h newsHandlers) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(201, gin.H{"okay": true})
	}
}

func (h newsHandlers) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(201, gin.H{"update": true})
	}
}

func (h newsHandlers) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"GetByID": true})
	}
}

func (h newsHandlers) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(201, gin.H{"Delete": true})
	}
}

func (h newsHandlers) GetNews() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"GetNews": true})
	}
}

func (h newsHandlers) SearchByTitle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"SearchByTitle": true})
	}
}
