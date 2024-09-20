package cmd

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/meiti-x/go-blog/config"
	"github.com/meiti-x/go-blog/internal/server"
	"github.com/meiti-x/go-blog/pkg/utils"
	"log"
)

func Execute() {
	// load configs
	configPath := utils.GetConfigPath("dev")

	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal("Could not read config file: ", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatal("Could not read config file: ", err)
	}

	// connect to postgres
	conn := fmt.Sprintf("user=%s dbname=%s sslmode=%s password=%s host=%s", cfg.Postgres.PostgresqlUser, cfg.Postgres.PostgresqlDbname, cfg.Postgres.PostgresqlSSLMode, cfg.Postgres.PostgresqlPassword, cfg.Postgres.PostgresqlHost)
	db, err := sqlx.Connect(cfg.Postgres.PgDriver, conn)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database connection established")
	}

	// create server
	s := server.NewServer(cfg)
	s.Run()
}
