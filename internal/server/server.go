package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/meiti-x/go-blog/config"
	"net/http"
)

type Server struct {
	server *config.ServerConfig
}

func NewServer(cfgFile *config.Config) *Server {
	return &Server{
		server: &cfgFile.Server,
	}

}

func (s *Server) Run() error {
	r := gin.Default()

	if err := s.MapNewsRoutes(r); err != nil {
		return err
	}
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "blad",
		})
	})

	r.Run(s.server.Port)

	fmt.Println("Server started on port " + "http://127.0.0.1" + s.server.Port)
	return nil
}
