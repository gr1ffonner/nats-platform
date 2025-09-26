# NATS Messaging Project

A Go project demonstrating NATS messaging patterns including pub/sub, load balancing, and JetStream persistence.

## Project Overview

This project showcases NATS messaging capabilities with:
- **Core NATS**: Pub/Sub messaging patterns
- **JetStream**: Persistent messaging with guaranteed delivery
- **Load Balancing**: Queue groups for distributed processing
- **Monitoring**: NATS server monitoring and metrics

## Services & Ports

| Service | Port | Description |
|---------|------|-------------|
| **NATS Server** | 4222 | Message broker |
| **NATS Monitor** | 8222 | NATS monitoring dashboard |

## Quick Start

### Prerequisites
- Docker & Docker Compose
- Go 1.21+
- Make

### Commands

#### Infrastructure Management
```bash
# Start NATS server
make up

# Stop NATS server
make down
```

#### Application
```bash
# Run consumer
make run-consumer

# Run producer
make run-producer
```

## Monitoring
- **NATS Monitor**: http://localhost:8222