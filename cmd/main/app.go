package main

import (
	"man_utd/pkg/logging"
	"man_utd/server"
)

func main() {
	logger := logging.GetLogger()

	srv := server.Server{}

	logger.Error("start server")
	if err := srv.Run("8000", nil); err != nil {
		logger.Fatal(err)
	}
}
