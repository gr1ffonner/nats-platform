package main

import (
	"context"
	"log/slog"
	"nats-platform/pkg/broker"
	"nats-platform/pkg/config"
	"nats-platform/pkg/logger"
)

const (
	subject = "test"
	message = "Hello, world"
)

type Message struct {
	Message string `json:"message"`
}

func main() {

	msg := Message{
		Message: "Hello, world",
	}

	// Load config first
	cfg, err := config.Load()
	if err != nil {
		slog.Error("Failed to load config", "error", err)
		panic(err)
	}
	slog.Info("Config loaded", "config", cfg)

	logger.InitLogger(cfg.Logger)
	logger := slog.Default()

	logger.Info("Config and logger initialized")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	natsClient, err := broker.NewNATS(ctx, cfg.NATS)
	if err != nil {
		slog.Error("Failed to create NATS client", "error", err)
		panic(err)
	}
	defer natsClient.Close()

	natsClient.PublishMessage(subject, message)
	natsClient.PublishMessage(subject, msg)

}
