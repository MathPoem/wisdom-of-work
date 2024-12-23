package main

import (
	"os"
	"os/signal"
	"syscall"
	"wisdom-of-work-client/pkg/client"
	"wisdom-of-work-client/internal/config"
	"wisdom-of-work-client/pkg/logger"
)

func main() {
	logger.InitLogger()
	log := logger.Log
	log.Info("Starting Wisdom of Work client...")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.WithError(err).Error("Failed to load config")
		os.Exit(1)
	}

	clientInstance := client.NewClient(cfg, log)
	go clientInstance.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down client gracefully...")
	clientInstance.Stop()
}
