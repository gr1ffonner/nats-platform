package broker

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"nats-platform/pkg/config"
	"time"

	"github.com/nats-io/nats.go"
)

type NATSClient struct {
	conn *nats.Conn
}

func NewNATS(ctx context.Context, cfg config.NATSConfig) (*NATSClient, error) {
	opts := []nats.Option{
		nats.Name(cfg.ClientName),
		nats.Timeout(10 * time.Second),
		nats.ReconnectWait(2 * time.Second),
		nats.MaxReconnects(5),
		nats.ReconnectJitter(100*time.Millisecond, 1*time.Second),
	}

	conn, err := nats.Connect(cfg.URL, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}

	// Test the connection
	if !conn.IsConnected() {
		conn.Close()
		return nil, fmt.Errorf("failed to establish NATS connection")
	}

	return &NATSClient{conn: conn}, nil
}

func (n *NATSClient) Close() {
	if n.conn != nil {
		n.conn.Close()
	}
}

// PublishOrder publishes an order to the specified subject
func (n *NATSClient) PublishMessage(subject string, message interface{}) error {
	slog.Info("Marshaling message", "subject", subject, "message", message)
	data, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	slog.Info("Publishing message", "subject", subject, "message", string(data))
	err = n.conn.Publish(subject, data)
	if err != nil {
		slog.Error("Failed to publish message", "subject", subject, "error", err)
		return fmt.Errorf("failed to publish message: %w", err)
	}

	slog.Info("Message published", "subject", subject, "message", string(data))

	return nil
}

// SubscribeToMessage subscribes to a message on the specified subject with a custom handler
// Pkg func for usage in microservices
func (n *NATSClient) SubscribeToMessage(subject string, handler func([]byte) error) (*nats.Subscription, error) {
	sub, err := n.conn.Subscribe(subject, func(msg *nats.Msg) {
		if err := handler(msg.Data); err != nil {
			slog.Error("Error processing message on subject", "subject", subject, "error", err)
			return
		}
		slog.Info("Message processed successfully", "subject", subject)
		msg.Ack()
	})
	if err != nil {
		slog.Error("Failed to subscribe to message", "subject", subject, "error", err)
		return nil, fmt.Errorf("failed to subscribe: %w", err)
	}

	return sub, nil
}

// ListenToMessage listens to a message on the specified subject with a default handler
// It will log the message and ack it
func (n *NATSClient) ListenToMessage(subject string) (*nats.Subscription, error) {
	slog.Info("Listening to message", "subject", subject)
	sub, err := n.conn.Subscribe(subject, func(msg *nats.Msg) {
		defer msg.Ack()

		// Try to parse as JSON for better logging
		var jsonMsg interface{}
		if err := json.Unmarshal(msg.Data, &jsonMsg); err == nil {
			slog.Info("Message received", "subject", subject, "message", jsonMsg)
		} else {
			// Fall back to string if not valid JSON
			slog.Warn("Message received is not valid JSON")
			slog.Info("Message received", "subject", subject, "message", string(msg.Data))
		}
	})
	if err != nil {
		slog.Error("Failed to subscribe to message", "subject", subject, "error", err)
		return nil, fmt.Errorf("failed to subscribe: %w", err)
	}

	return sub, nil
}
