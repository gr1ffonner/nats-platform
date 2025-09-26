CMD_DIR= ./cmd/


# Run application with PostgreSQL
run-consumer:
	@export $$(grep -v '^#' .env | xargs) >/dev/null 2>&1; \
	go run $(CMD_DIR)/consumer/main.go

run-producer:
	@export $$(grep -v '^#' .env | xargs) >/dev/null 2>&1; \
	go run $(CMD_DIR)/producer/main.go


# Start all services
up:
	COMPOSE_PROJECT_NAME=go-platform docker compose -f docker-compose.yml --profile=test --env-file=.env-docker up -d --build

# Start only infrastructure
up-infra:
	COMPOSE_PROJECT_NAME=go-platform docker compose -f docker-compose.yml --env-file=.env up -d 

# Stop all services
down:
	COMPOSE_PROJECT_NAME=go-platform docker compose -f docker-compose.yml --profile=test --env-file=.env-docker down --remove-orphans

# Stop and remove volumes
down-infra:
	COMPOSE_PROJECT_NAME=go-platform docker compose -f docker-compose.yml --env-file=.env down --remove-orphans

clean:
	COMPOSE_PROJECT_NAME=go-platform docker compose -f docker-compose.yml --profile=test --env-file=.env-docker down -v --remove-orphans

clean-infra:
	COMPOSE_PROJECT_NAME=go-platform docker compose -f docker-compose.yml --env-file=.env down -v --remove-orphans
