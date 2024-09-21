package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	models "github.com/meiti-x/go-blog/internal/_models"
	"github.com/meiti-x/go-blog/internal/blog"
)

type newsHandlers struct {
	db *sqlx.DB
	uc blog.UseCase
}

func NewBlogsHandlers(db *sqlx.DB, uc blog.UseCase) blog.Handlers {
	return &newsHandlers{
		db: db,
		uc: uc,
	}
}

func (h newsHandlers) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var item *models.Posts

		if err := c.ShouldBindJSON(&item); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		item, err := h.uc.Create(c, item)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, gin.H{"data": item})
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
