package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/meiti-x/go-blog/config"
	"log"
	"net/http"
)

type Server struct {
	server *config.ServerConfig
	db     *sqlx.DB
}

func NewServer(cfgFile *config.Config, db *sqlx.DB) *Server {
	return &Server{
		server: &cfgFile.Server,
		db:     db,
	}

}

func (s *Server) Run() error {
	r := gin.Default()

	if err := s.MapRoutes(r, s.db); err != nil {
		return err
	}
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "blad",
		})
	})

	err := r.Run(s.server.Port)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server started on port " + "http://127.0.0.1" + s.server.Port)
	return nil
}
