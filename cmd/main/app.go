package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"man_utd/internal/config"
	"man_utd/internal/handler"
	"man_utd/internal/repository"
	"man_utd/internal/service"
	"man_utd/pkg/client"
	"man_utd/pkg/logging"
	"man_utd/server"
	"os"
)

func main() {
	logger := logging.GetLogger()
	cfg := config.GetConfig()
	srv := server.Server{}

	if err := godotenv.Load(); err != nil {
		logger.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := client.NewPostgresDB(&client.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: cfg.DB.Username,
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   cfg.DB.DBName,
		SSLMode:  cfg.DB.SSLMode,
	})

	if err != nil {
		logger.Fatalf("error creating db postgres: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	logger.Error("start server")
	if err := srv.Run(cfg.Port, handlers.InitRoutes()); err != nil {
		logger.Fatal(err)
	}
}
