package server

import (
	"github.com/gin-gonic/gin"
	"github.com/meiti-x/go-blog/config"
	"net/http"
)

type Server struct {
	server *config.ServerConfig
}

func NewServer() *Server {
	return &Server{}

}

func (s *Server) Run() error {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "blad",
		})
	})
	r.Run()
	return nil
}
