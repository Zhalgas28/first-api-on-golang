package main

import (
	_ "github.com/lib/pq"
	"man_utd/internal/config"
	"man_utd/pkg/logging"
	"man_utd/server"
)

func main() {
	logger := logging.GetLogger()
	cfg := config.GetConfig()
	srv := server.Server{}

	logger.Error("start server")
	if err := srv.Run(cfg.Port, nil); err != nil {
		logger.Fatal(err)
	}
}
