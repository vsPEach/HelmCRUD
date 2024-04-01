COMPOSE_PATH=crud-service/deploy/docker-compose.yml

up: 
	docker compose -f $(COMPOSE_PATH) up -d --build --force-recreate

down:
	docker compose -f $(COMPOSE_PATH) down

.PHONY: up down
