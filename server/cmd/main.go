package main

import (
	"os"
	"os/signal"
	"syscall"

	"wisdom-of-work-server/internal/config"
	"wisdom-of-work-server/pkg/logger"
	"wisdom-of-work-server/pkg/tcp_server"
)

func main() {

	logger.InitLogger()
	log := logger.Log
	log.Info("Starting Wisdom of Work server...")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.WithError(err).Error("Failed to load config")
		os.Exit(1)
	}

	server, err := tcp_server.NewServer(cfg, log)
	if err != nil {
		log.WithError(err).Error("Failed to create server")
		os.Exit(1)
	}
	go server.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down server gracefully...")
	server.Stop()
}
