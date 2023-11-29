package main

import (
	"log"
	"github.com/furkanansn/fasttrack-task/api"
)

func main() {
	cfg := api.ServerConfig{
		Port: 8080,
	}

	server, err := api.NewServer(cfg)
	if err != nil {
		log.Fatalf("Error creating server: %v", err)
	}

	api.InitServer(server)

	server.SetupRoutes()
	server.Start()
}
