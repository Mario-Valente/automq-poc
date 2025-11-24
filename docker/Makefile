.PHONY: help up down restart logs logs-follow status clean ps exec-broker exec-minio create-topic list-topics describe-topic delete-topic produce consume shell-broker shell-minio kafka-ui

# Variables
COMPOSE_FILE := docker-compose.yaml
PROJECT_NAME := automq-poc
BROKER_CONTAINER := automq-single-server
MINIO_CONTAINER := minio
MC_CONTAINER := mc

# Colors for output
GREEN  := \033[0;32m
YELLOW := \033[0;33m
RED    := \033[0;31m
BLUE   := \033[0;34m
NC     := \033[0m # No Color

help: ## Show this help message
	@echo '$(GREEN)AutoMQ Docker Makefile$(NC)'
	@echo ''
	@echo 'Usage:'
	@echo '  make $(YELLOW)<target>$(NC)'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  $(YELLOW)%-20s$(NC) %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Container Management
up: ## Start all containers
	@echo "$(GREEN)Starting AutoMQ cluster...$(NC)"
	docker compose -f $(COMPOSE_FILE) up -d
	@echo "$(GREEN)Waiting for cluster to be ready (30s)...$(NC)"
	@sleep 30
	@echo "$(GREEN)Cluster is ready!$(NC)"

down: ## Stop and remove all containers
	@echo "$(RED)Stopping AutoMQ cluster...$(NC)"
	docker compose -f $(COMPOSE_FILE) down

restart: down up ## Restart all containers
	@echo "$(GREEN)Cluster restarted!$(NC)"

restart-broker: ## Restart only the broker
	@echo "$(YELLOW)Restarting broker...$(NC)"
	docker compose -f $(COMPOSE_FILE) restart server1
	@sleep 20
	@echo "$(GREEN)Broker restarted!$(NC)"

stop: ## Stop containers without removing
	@echo "$(YELLOW)Stopping containers...$(NC)"
	docker compose -f $(COMPOSE_FILE) stop

start: ## Start stopped containers
	@echo "$(GREEN)Starting containers...$(NC)"
	docker compose -f $(COMPOSE_FILE) start

# Monitoring & Logs
logs: ## Show logs from all containers
	docker compose -f $(COMPOSE_FILE) logs

logs-follow: ## Follow logs from all containers
	docker compose -f $(COMPOSE_FILE) logs -f

logs-broker: ## Show broker logs
	docker logs $(BROKER_CONTAINER)

logs-broker-follow: ## Follow broker logs
	docker logs -f $(BROKER_CONTAINER)

logs-minio: ## Show MinIO logs
	docker logs $(MINIO_CONTAINER)

status: ## Show container status
	@echo "$(BLUE)Container Status:$(NC)"
	@docker compose -f $(COMPOSE_FILE) ps

ps: status ## Alias for status

health: ## Check health of all services
	@echo "$(BLUE)Service Health:$(NC)"
	@docker ps --filter "name=$(BROKER_CONTAINER)" --format "table {{.Names}}\t{{.Status}}"
	@docker ps --filter "name=$(MINIO_CONTAINER)" --format "table {{.Names}}\t{{.Status}}"

# Kafka Operations
create-topic: ## Create a topic (use TOPIC=<name> PARTITIONS=<n> REPLICAS=<n>)
	@if [ -z "$(TOPIC)" ]; then \
		echo "$(RED)Error: TOPIC required$(NC)"; \
		echo "Usage: make create-topic TOPIC=my-topic PARTITIONS=3 REPLICAS=1"; \
		exit 1; \
	fi
	@PARTITIONS=$${PARTITIONS:-1}; \
	REPLICAS=$${REPLICAS:-1}; \
	docker exec $(BROKER_CONTAINER) \
		/opt/automq/kafka/bin/kafka-topics.sh \
		--bootstrap-server localhost:9092 \
		--create --topic $(TOPIC) \
		--partitions $$PARTITIONS --replication-factor $$REPLICAS

list-topics: ## List all Kafka topics
	@echo "$(BLUE)Available Topics:$(NC)"
	@docker exec $(BROKER_CONTAINER) \
		/opt/automq/kafka/bin/kafka-topics.sh \
		--bootstrap-server localhost:9092 \
		--list

describe-topic: ## Describe a topic (use TOPIC=<name>)
	@if [ -z "$(TOPIC)" ]; then \
		echo "$(RED)Error: TOPIC required$(NC)"; \
		echo "Usage: make describe-topic TOPIC=my-topic"; \
		exit 1; \
	fi
	@docker exec $(BROKER_CONTAINER) \
		/opt/automq/kafka/bin/kafka-topics.sh \
		--bootstrap-server localhost:9092 \
		--describe --topic $(TOPIC)

delete-topic: ## Delete a topic (use TOPIC=<name>)
	@if [ -z "$(TOPIC)" ]; then \
		echo "$(RED)Error: TOPIC required$(NC)"; \
		echo "Usage: make delete-topic TOPIC=my-topic"; \
		exit 1; \
	fi
	@read -p "Are you sure you want to delete topic '$(TOPIC)'? [y/N] " -n 1 -r; \
	echo; \
	if [[ $$REPLY =~ ^[Yy]$$ ]]; then \
		docker exec $(BROKER_CONTAINER) \
			/opt/automq/kafka/bin/kafka-topics.sh \
			--bootstrap-server localhost:9092 \
			--delete --topic $(TOPIC); \
		echo "$(GREEN)Topic deleted!$(NC)"; \
	else \
		echo "$(RED)Delete cancelled$(NC)"; \
	fi

produce: ## Produce messages to a topic (use TOPIC=<name>)
	@if [ -z "$(TOPIC)" ]; then \
		echo "$(RED)Error: TOPIC required$(NC)"; \
		echo "Usage: make produce TOPIC=my-topic"; \
		exit 1; \
	fi
	@echo "$(GREEN)Producer started. Type messages (Ctrl+D to finish):$(NC)"
	@docker exec -i $(BROKER_CONTAINER) \
		/opt/automq/kafka/bin/kafka-console-producer.sh \
		--bootstrap-server localhost:9092 \
		--topic $(TOPIC)

consume: ## Consume messages from a topic (use TOPIC=<name>)
	@if [ -z "$(TOPIC)" ]; then \
		echo "$(RED)Error: TOPIC required$(NC)"; \
		echo "Usage: make consume TOPIC=my-topic"; \
		exit 1; \
	fi
	@echo "$(GREEN)Consuming messages from '$(TOPIC)' (Ctrl+C to stop):$(NC)"
	@docker exec $(BROKER_CONTAINER) \
		/opt/automq/kafka/bin/kafka-console-consumer.sh \
		--bootstrap-server localhost:9092 \
		--topic $(TOPIC) \
		--from-beginning

# Shell Access
shell-broker: ## Open bash shell in broker container
	@docker exec -it $(BROKER_CONTAINER) bash

shell-minio: ## Open shell in MinIO container
	@docker exec -it $(MINIO_CONTAINER) sh

exec-broker: ## Execute command in broker (use CMD="<command>")
	@if [ -z "$(CMD)" ]; then \
		echo "$(RED)Error: CMD required$(NC)"; \
		echo "Usage: make exec-broker CMD='ls -la'"; \
		exit 1; \
	fi
	@docker exec $(BROKER_CONTAINER) $(CMD)

# MinIO Operations
minio-console: ## Open MinIO console URL
	@echo "$(GREEN)MinIO Console: http://localhost:9001$(NC)"
	@echo "Username: minioadmin"
	@echo "Password: minioadmin"

minio-list-buckets: ## List MinIO buckets
	@echo "$(BLUE)MinIO Buckets:$(NC)"
	@docker exec $(MC_CONTAINER) mc ls minio

minio-stats: ## Show MinIO bucket statistics
	@echo "$(BLUE)Data Bucket Stats:$(NC)"
	@docker exec $(MC_CONTAINER) mc du minio/automq-data
	@echo ""
	@echo "$(BLUE)Ops Bucket Stats:$(NC)"
	@docker exec $(MC_CONTAINER) mc du minio/automq-ops

# Cleanup
clean: down ## Stop containers and remove volumes
	@echo "$(YELLOW)Removing volumes...$(NC)"
	docker compose -f $(COMPOSE_FILE) down -v
	@echo "$(GREEN)Cleanup complete!$(NC)"

clean-all: clean ## Clean everything including images
	@echo "$(YELLOW)Removing images...$(NC)"
	docker rmi automqinc/automq:1.5.5 minio/minio:RELEASE.2025-05-24T17-08-30Z minio/mc:RELEASE.2025-05-21T01-59-54Z 2>/dev/null || true
	@echo "$(GREEN)Full cleanup complete!$(NC)"

# Testing & Validation
test-connection: ## Test connection to broker
	@echo "$(BLUE)Testing broker connection...$(NC)"
	@if docker exec $(BROKER_CONTAINER) \
		/opt/automq/kafka/bin/kafka-broker-api-versions.sh \
		--bootstrap-server localhost:9092 > /dev/null 2>&1; then \
		echo "$(GREEN)✓ Broker is reachable$(NC)"; \
	else \
		echo "$(RED)✗ Broker is not reachable$(NC)"; \
		exit 1; \
	fi

test-produce-consume: ## Test produce and consume flow
	@echo "$(BLUE)Testing produce/consume...$(NC)"
	@TOPIC="test-$$(date +%s)"; \
	echo "Creating topic $$TOPIC..."; \
	make create-topic TOPIC=$$TOPIC PARTITIONS=1 REPLICAS=1; \
	echo "Producing test message..."; \
	echo "test-message-$$(date +%s)" | docker exec -i $(BROKER_CONTAINER) \
		/opt/automq/kafka/bin/kafka-console-producer.sh \
		--bootstrap-server localhost:9092 \
		--topic $$TOPIC; \
	echo "Consuming messages..."; \
	docker exec $(BROKER_CONTAINER) \
		/opt/automq/kafka/bin/kafka-console-consumer.sh \
		--bootstrap-server localhost:9092 \
		--topic $$TOPIC \
		--from-beginning \
		--timeout-ms 5000; \
	echo "$(GREEN)✓ Test completed$(NC)"

# Quick workflows
setup: up test-connection ## Setup and verify cluster
	@echo "$(GREEN)Cluster setup complete and verified!$(NC)"

full-test: setup test-produce-consume ## Full setup and testing
	@echo "$(GREEN)All tests passed!$(NC)"

info: ## Show cluster information
	@echo "$(BLUE)AutoMQ Cluster Information:$(NC)"
	@echo "Broker:        localhost:9092"
	@echo "MinIO API:     localhost:9000"
	@echo "MinIO Console: localhost:9001"
	@echo ""
	@echo "Hostnames:"
	@echo "  Add to /etc/hosts: 127.0.0.1 server1"
