package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/meiti-x/go-blog/internal/blog"
)

type newsHandlers struct {
	db *sqlx.DB
}

func NewBlogsHandlers(db *sqlx.DB) blog.Handlers {
	return &newsHandlers{
		db: db,
	}
}

func (h newsHandlers) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newBlog struct {
			Title   string `json:"title" binding:"required"`
			Content string `json:"content" binding:"required"`
		}

		if err := c.ShouldBindJSON(&newBlog); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		fmt.Println(newBlog)
		query := `INSERT INTO blogs (title, content) VALUES ($1, $2)`
		_, err := h.db.Exec(query, newBlog.Title, newBlog.Content)
		if err != nil {
			c.JSON(500, gin.H{"error": err})
			return
		}

		c.JSON(201, gin.H{"message": "Blog created successfully"})
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
