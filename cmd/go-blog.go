package cmd

import (
	"fmt"
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
	fmt.Println(cfgFile)
	s := server.NewServer()
	s.Run()
}
