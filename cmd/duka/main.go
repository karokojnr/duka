package main

import (
	"github.com/karokojnr/duka/config"
	"github.com/karokojnr/duka/internal/api"
	"log"
)

func run() error {
	cfg, err := config.SetUpEnvironment()
	if err != nil {
		log.Fatalf("configurations not loaded %v\n", err)
	}
	api.StartServer(cfg)
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
