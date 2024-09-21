package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/meiti-x/go-blog/config"
	"log"
)

type Server struct {
	cfg *config.Config
	db  *sqlx.DB
}

func NewServer(cfgFile *config.Config, db *sqlx.DB) *Server {
	return &Server{
		cfg: cfgFile,
		db:  db,
	}

}

func (s *Server) Run() error {
	r := gin.Default()

	if err := s.MapRoutes(r, s.db, s.cfg); err != nil {
		return err
	}

	err := r.Run(s.cfg.Server.Port)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server started on port " + "http://127.0.0.1" + s.cfg.Server.Port)
	return nil
}
