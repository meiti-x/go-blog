package cmd

import (
	"github.com/meiti-x/go-blog/config"
	"github.com/meiti-x/go-blog/internal/server"
	"github.com/meiti-x/go-blog/pkg/utils"
	"log"
)

func Execute() {
	configPath := utils.GetConfigPath("dev")

	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal("Could not read config file: ", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatal("Could not read config file: ", err)
	}

	s := server.NewServer(cfg)
	s.Run()
}
