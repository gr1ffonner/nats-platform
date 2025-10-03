package main

import (
	"context"
	"log/slog"
	"nats-platform/pkg/broker"
	"nats-platform/pkg/config"
	"nats-platform/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

const (
	subject = "test"
)

func main() {
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

	notifyCtx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	natsClient, err := broker.NewNATS(ctx, cfg.NATS)
	if err != nil {
		slog.Error("Failed to create NATS client", "error", err)
		panic(err)
	}
	defer natsClient.Close()

	sub, err := natsClient.ListenToMessage(subject)
	if err != nil {
		slog.Error("Failed to listen to message", "error", err)
		panic(err)
	}
	defer sub.Unsubscribe()

	<-notifyCtx.Done()
	slog.Info("Shutting down consumer")

}
