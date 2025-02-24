package main

import (
	"fmt"
	"github.com/karokojnr/duka/config"
	"github.com/karokojnr/duka/internal/api"
	"log"
)

func run() error {
	cfg, err := config.SetUpEnvironment()
	if err != nil {
		return fmt.Errorf("config.SetupEnvironment: %w", err)
	}
	api.StartServer(cfg)
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
