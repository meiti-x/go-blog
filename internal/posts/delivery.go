package posts

import "github.com/gin-gonic/gin"

// Posts HTTP Handlers interface
type Handlers interface {
	Create() gin.HandlerFunc
	Update() gin.HandlerFunc
	GetByID() gin.HandlerFunc
	Delete() gin.HandlerFunc
	GetNews() gin.HandlerFunc
	SearchByTitle() gin.HandlerFunc
}
