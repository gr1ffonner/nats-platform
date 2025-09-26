# NATS JetStream Documentation

## Overview

NATS JetStream is the persistence layer built on top of core NATS. It provides message persistence, guaranteed delivery, and advanced streaming capabilities.

## JetStream Features

### Core Capabilities
- **Message Persistence**: Messages stored on disk
- **At-Least-Once Delivery**: Guaranteed message delivery
- **Streams**: Persistent message storage
- **Consumers**: Durable subscriptions
- **Replay**: Ability to replay messages from any point
- **Durability**: Survives server restarts

### JetStream vs Core NATS

| Feature | Core NATS | JetStream |
|---------|-----------|-----------|
| Persistence | No | Yes |
| Delivery Guarantee | At-Most-Once | At-Least-Once |
| Replay | No | Yes |
| Durability | No | Yes |
| Message Retention | No | Configurable |
| Dead Letter Queues | No | Yes |
