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
	data, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	err = n.conn.Publish(subject, data)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}

// SubscribeToOrders subscribes to orders on the specified subject
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
